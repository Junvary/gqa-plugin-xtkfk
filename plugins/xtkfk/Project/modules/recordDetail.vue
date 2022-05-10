<template>
    <q-dialog v-model="recordDetailVisible" position="top">
        <q-card style="width: 1400px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{ formTypeName }}项目:
                    {{ recordDetail.value.title }}
                </div>
            </q-card-section>

            <q-separator />

            <q-card-section>
                <q-form ref="recordDetailForm">
                    <gqa-form-top :recordDetail="recordDetail"></gqa-form-top>
                    <div class="row">
                        <q-input class="col" v-model="recordDetail.value.project_name" label="项目名称"
                            :rules="[val => val && val.length > 0 || '必须输入项目名称']" />
                        <q-field dense label="是否启用" stack-label>
                            <template v-slot:control>
                                <q-option-group v-model="recordDetail.value.status" :options="dictOptions.statusOnOff"
                                    color="primary" inline>
                                </q-option-group>
                            </template>
                        </q-field>
                    </div>
                    <div class="row">
                        <q-input class="col" v-model="recordDetail.value.demand" label="需求单位" />
                        <q-select class="col" v-model="recordDetail.value.language" :options="dictOptions.codeLanguage"
                            multiple clearable emit-value map-options
                            :rules="[val => val && val.length > 0 || '必须选择项目语言']" label="项目语言" />
                        <q-select class="col" v-model="recordDetail.value.node"
                            :options="dictOptions.projectNode.filter(item => item.dict_code.indexOf('p') !== -1)"
                            clearable emit-value map-options :rules="[val => val && val.length > 0 || '必须选择项目节点']"
                            label="项目节点" />
                        <GqaSeleteUser className="col" label="牵头人" v-model:selectUser="recordDetail.value.leader"
                            v-model:selectUsername="recordDetail.value.leader_username" selection="single" />
                    </div>
                    <q-select bottom-slots label="参与人" v-model="recordDetail.value.player" use-input use-chips multiple
                        hide-dropdown-icon input-debounce="0" new-value-mode="add-unique">
                        <template v-slot:hint>
                            使用回车确定单个人选
                        </template>
                    </q-select>
                    <q-input class="col" v-model="recordDetail.value.memo" type="textarea" label="备注" />
                </q-form>
            </q-card-section>

            <q-separator />

            <q-card-actions align="right">
                <q-btn :label="$t('Save') + formTypeName" color="primary" @click="handleAddOrEidt" />
                <q-btn :label="$t('Cancel')" color="negative" v-close-popup />
            </q-card-actions>

            <q-inner-loading :showing="loading">
                <q-spinner-gears size="50px" color="primary" />
            </q-inner-loading>
        </q-card>
    </q-dialog>
</template>


<script setup>
import useRecordDetail from 'src/composables/useRecordDetail'
import GqaAvatar from 'src/components/GqaAvatar'
import GqaShowName from 'src/components/GqaShowName'
import { postAction } from 'src/api/manage'
import { useStorageStore } from 'src/stores/storage'
import { ref, computed } from 'vue'
import { useQuasar } from 'quasar'
import { useI18n } from 'vue-i18n'
import GqaSeleteUser from 'src/components/GqaSeleteUser'

const $q = useQuasar()
const { t } = useI18n()
const storageStore = useStorageStore()
const gqaBackend = computed(() => storageStore.GetGqaBackend())
const emit = defineEmits(['handleFinish'])
const url = {
    add: 'plugin-xtkfk/add-project',
    edit: 'plugin-xtkfk/edit-project',
    queryById: 'plugin-xtkfk/query-project-by-id',
}
const {
    dictOptions,
    showDateTime,
    formType,
    formTypeName,
    recordDetail,
    recordDetailVisible,
    loading,
    // show,
    recordDetailForm,
    // handleAddOrEidt
} = useRecordDetail(url, emit)

const show = (row) => {
    loading.value = true
    recordDetail.value = {}
    recordDetailVisible.value = true
    if (row && row.id) {
        handleQueryById(row.id)
    } else {
        recordDetail.value = {}
        recordDetailVisible.value = true
        loading.value = false
    }
}

defineExpose({
    show,
    formType
})



const handleQueryById = async (id) => {
    const res = await postAction(url.queryById, { id, })
    if (res.code === 1) {
        recordDetail.value = res.data.records
        recordDetail.value.player = recordDetail.value.player !== '' ? recordDetail.value.player.split(',') : []
        recordDetail.value.language = recordDetail.value.language.split(',')
    }
    loading.value = false
}

const handleAddOrEidt = async () => {
    const success = await recordDetailForm.value.validate()
    if (success) {
        recordDetail.value.player = recordDetail.value.player.join(',')
        recordDetail.value.language = recordDetail.value.language.join(',')

        if (formType.value === 'edit') {
            if (url === undefined || !url.edit) {
                $q.notify({
                    type: 'negative',
                    message: '请先配置url',
                })
                return
            }
            const res = await postAction(url.edit, recordDetail.value)
            if (res.code === 1) {
                $q.notify({
                    type: 'positive',
                    message: res.message,
                })
                recordDetailVisible.value = false
            }
        } else if (formType.value === 'add') {
            if (url === undefined || !url.add) {
                $q.notify({
                    type: 'negative',
                    message: '请先配置url',
                })
                return
            }
            const res = await postAction(url.add, recordDetail.value)
            if (res.code === 1) {
                $q.notify({
                    type: 'positive',
                    message: res.message,
                })
                recordDetailVisible.value = false
            }
        } else {
            $q.notify({
                type: 'negative',
                message: '无法新增或编辑！',
            })
        }
        emit('handleFinish')
    } else {
        $q.notify({
            type: 'negative',
            message: '请完善表格信息！',
        })
    }
}
</script>
