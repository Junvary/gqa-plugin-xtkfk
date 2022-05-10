<template>
    <q-dialog v-model="recordDetailVisible" position="top">
        <q-card style="width: 1400px; max-width: 80vw;">
            <q-card-section>
                <div class="text-h6">
                    {{ formTypeName }}荣誉认证:
                    {{ recordDetail.value.title }}
                </div>
            </q-card-section>

            <q-separator />

            <q-card-section>
                <q-form ref="recordDetailForm">
                    <gqa-form-top :recordDetail="recordDetail"></gqa-form-top>
                    <div class="row">
                        <q-input class="col" v-model="recordDetail.value.title" label="标题"
                            :rules="[val => val && val.length > 0 || '必须输入标题']" />
                        <q-field dense label="是否启用" stack-label>
                            <template v-slot:control>
                                <q-option-group v-model="recordDetail.value.status" :options="dictOptions.statusOnOff"
                                    color="primary" inline>
                                </q-option-group>
                            </template>
                        </q-field>
                    </div>
                    <div class="row ">
                        <GqaUpload class="col" v-model:attachment="recordDetail.value.attachment" />
                        <q-separator vertical spaced />
                        <q-input class="col" v-model="recordDetail.value.memo" type="textarea" label="备注" />
                    </div>
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
import GqaEditor from 'src/components/GqaEditor'
import GqaUpload from 'src/components/GqaUpload'

const $q = useQuasar()
const { t } = useI18n()
const storageStore = useStorageStore()
const gqaBackend = computed(() => storageStore.GetGqaBackend())
const emit = defineEmits(['handleFinish'])
const url = {
    add: 'plugin-xtkfk/add-honour',
    edit: 'plugin-xtkfk/edit-honour',
    queryById: 'plugin-xtkfk/query-honour-by-id',
}
const {
    dictOptions,
    showDateTime,
    formType,
    formTypeName,
    recordDetail,
    recordDetailVisible,
    loading,
    show,
    recordDetailForm,
    handleAddOrEidt
} = useRecordDetail(url, emit)

defineExpose({
    show,
    formType
})
</script>
