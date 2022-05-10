<template>
    <q-dialog v-model="projectDetailVisible" position="top">
        <q-card style="width: 1400px; max-width: 80vw;">
            <q-card-section>
                <div class="row text-h6">
                    {{ project.project_name }}
                    <q-space></q-space>
                    <q-btn label="保存" color="primary" @click="handleSave" />
                </div>
            </q-card-section>

            <q-separator />
            <q-scroll-area style="height: 85vh">
                <q-card-section>
                    <q-form ref="projectDetailForm">
                        <q-timeline layout="comfortable" side="right" color="primary">
                            <q-timeline-entry class="project-style" :icon="checkIcon(item)" :color="checkColor(item)"
                                v-for="(item, index) in dictOptions.projectNode.filter(item => item.dict_code.indexOf('p') !== -1)"
                                :key="index" :side="side[index]">
                                <template v-slot:title>
                                    {{ item.dict_label }}
                                    <span
                                        v-if="projectForm[item.dict_code].timeRange && projectForm[item.dict_code].timeRange.from"
                                        style="float: right">
                                        {{ projectForm[item.dict_code].timeRange.from }}
                                        至
                                        {{ projectForm[item.dict_code].timeRange.to }}
                                    </span>

                                </template>
                                <template v-slot:subtitle>
                                    <q-date v-model="projectForm[item.dict_code].timeRange" range
                                        :color="checkColor(item)" mask="YYYY-MM-DD" />
                                    <q-separator spaced />
                                </template>
                                <div>
                                    <span class="text-subtitle5">
                                        {{ item.remark }}
                                    </span>
                                    <q-input outlined v-model="projectForm[item.dict_code].content" autogrow
                                        :label="item.dict_label + ' 备注'" />
                                    <GqaUpload v-model:attachment="projectForm[item.dict_code].attachment"
                                        :color="checkColor(item)" />
                                </div>
                            </q-timeline-entry>

                        </q-timeline>

                        <q-separator spaced />

                        <q-card-section>
                            <div class="row text-h6 justify-center">
                                其他材料
                            </div>
                        </q-card-section>

                        <q-timeline layout="comfortable" side="right" color="primary">
                            <q-timeline-entry class="project-style" :icon="checkIcon(item)" :color="checkColor(item)"
                                v-for="(item, index) in dictOptions.projectNode.filter(item => item.dict_code.indexOf('m') !== -1)"
                                :key="index" :title="item.dict_label" :side="side[index]">
                                <template v-slot:subtitle>
                                    <q-date v-model="projectForm[item.dict_code].timeRange" range
                                        :color="checkColor(item)" mask="YYYY-MM-DD" />
                                    <q-separator spaced />
                                </template>
                                <div>
                                    <span class="text-subtitle5">
                                        {{ item.remark }}
                                    </span>
                                    <q-input outlined v-model="projectForm[item.dict_code].content" autogrow
                                        :label="item.dict_label + ' 备注'" />
                                    <GqaUpload title="上传附件" v-model:attachment="projectForm[item.dict_code].attachment"
                                        :color="checkColor(item)" />
                                </div>
                            </q-timeline-entry>
                        </q-timeline>
                    </q-form>
                </q-card-section>
            </q-scroll-area>


            <q-inner-loading :showing="loading">
                <q-spinner-gears size="50px" color="primary" />
            </q-inner-loading>
        </q-card>
    </q-dialog>
</template>

<script setup>
import GqaUpload from 'src/components/GqaUpload'
import { computed, ref } from 'vue';
import { postAction } from 'src/api/manage'
import { DictOptions } from 'src/utils/dict'
import { useQuasar } from 'quasar';
import { useI18n } from 'vue-i18n';

const $q = useQuasar();
const { t } = useI18n();

const checkColor = computed(() => {
    return (item) => {
        if (projectForm.value[item.dict_code].content || projectForm.value[item.dict_code].attachment || (projectForm.value[item.dict_code].timeRange && projectForm.value[item.dict_code].timeRange.from && projectForm.value[item.dict_code].timeRange.to)) {
            return 'primary'
        } else {
            return 'grey-5'
        }
    }
})
const checkIcon = computed(() => {
    return (item) => {
        if (projectForm.value[item.dict_code].content || projectForm.value[item.dict_code].attachment || (projectForm.value[item.dict_code].timeRange && projectForm.value[item.dict_code].timeRange.from && projectForm.value[item.dict_code].timeRange.to)) {
            return 'check'
        } else {
            return 'question_mark'
        }
    }
})
const loading = ref(false)
const dictOptions = ref({})
const side = ref([])
const projectDetailVisible = ref(false)
const project = ref({})
const projectForm = ref({})
const url = {
    add: 'plugin-xk/add-project',
    edit: 'plugin-xk/edit-project-detail',
    queryById: 'plugin-xk/query-project-by-id',
}

const show = async (row) => {
    loading.value = true
    projectForm.value = {}
    side.value = ['left', 'right']
    project.value = row
    dictOptions.value = await DictOptions()
    const projectNode = dictOptions.value.projectNode
    for (let item of projectNode) {
        projectForm.value[item.dict_code] = {
            content: '',
            attachment: '',
            timeRange: {
                from: '',
                to: '',
            },
        }
        side.value = side.value.concat(side.value)
    }
    projectDetailVisible.value = true
    handleQueryById(project.value.id)
}
defineExpose({
    show,
})
const handleQueryById = async (id) => {
    const res = await postAction(url.queryById, { id, })
    if (res.code === 1) {
        const projectDetail = res.data.records.project_detail
        for (let item of projectDetail) {
            projectForm.value[item.project_node].content = item.content
            projectForm.value[item.project_node].attachment = item.attachment
            projectForm.value[item.project_node].timeRange = Object.assign({}, projectForm.value[item.project_node].timeRange, {
                from: item.start_time,
                to: item.end_time,
            })
        }
    }
    loading.value = false
}
const projectDetailForm = ref(null)
const handleSave = async () => {
    const success = await projectDetailForm.value.validate()
    if (success) {
        const projectDetail = []
        for (let key in projectForm.value) {
            projectDetail.push({
                project_id: project.value.project_id,
                project_node: key,
                content: projectForm.value[key].content,
                attachment: projectForm.value[key].attachment,
                start_time: projectForm.value[key].timeRange.from,
                end_time: projectForm.value[key].timeRange.to,
            })
        }
        postAction(url.edit, {
            project_id: project.value.project_id,
            project_detail: projectDetail,
        }).then((res) => {
            if (res.code === 1) {
                $q.notify({
                    type: 'positive',
                    message: res.message,
                })
                project.valueDetailVisible = false
            }
        })
    } else {
        $q.notify({
            type: 'negative',
            message: t('FixForm'),
        })
    }
}

</script>

<style lang="scss" scoped>
::v-deep(.q-timeline__subtitle) {
    opacity: 1;
}
</style>