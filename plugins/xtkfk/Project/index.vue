<template>
    <q-page padding>

        <div class="row q-gutter-md items-center" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.project_name" label="项目名称" />
            <q-btn color="primary" @click="handleSearch" label="搜索" />
            <q-btn color="primary" @click="resetSearch" label="重置" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn color="primary" @click="showAddForm()" label="新增项目" />
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

            <template v-slot:body-cell-leader="props">
                <q-td :props="props">
                    <GqaShowName v-if="props.row.leader" :customNameObject="props.row.leader" />
                </q-td>
            </template>

            <template v-slot:body-cell-language="props">
                <q-td :props="props">
                    <GqaDictShow dictName="codeLanguage" :dictCode="props.row.language"
                        v-if="props.row.language !== ''" />
                </q-td>
            </template>

            <template v-slot:body-cell-node="props">
                <q-td :props="props">
                    <GqaDictShow dictName="projectNode" :dictCode="props.row.node" v-if="props.row.node !== ''" />
                </q-td>
            </template>

            <template v-slot:body-cell-status="props">
                <q-td :props="props">
                    <GqaDictShow dictName="statusOnOff" :dictCode="props.row.status" />
                </q-td>
            </template>

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="warning" @click="showProjectDetail(props.row)" label="维护进度" />
                        <q-btn color="primary" @click="showEditForm(props.row)" label="编辑" />
                        <q-btn color="negative" @click="handleDelete(props.row)" label="删除" />
                    </div>
                </q-td>
            </template>
        </q-table>
        <recordDetail ref="recordDetailDialog" @handleFinish="handleFinish" />
        <ProjectDetailDialog ref="projectDetailDialog" />
    </q-page>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { useQuasar } from 'quasar'
import { postAction } from 'src/api/manage'
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import recordDetail from './modules/recordDetail'
import ProjectDetailDialog from './modules/ProjectDetailDialog'
import GqaShowName from 'src/components/GqaShowName'

const $q = useQuasar()
const { t } = useI18n()
const url = {
    list: 'plugin-xtkfk/get-project-list',
    delete: 'plugin-xtkfk/delete-project',
}
const columns = computed(() => {
    return [
        { name: 'project_name', align: 'center', label: '项目名称', field: 'project_name' },
        { name: 'demand', align: 'center', label: '需求单位', field: 'demand' },
        { name: 'leader', align: 'center', label: '牵头人', field: 'leader' },
        { name: 'player', align: 'center', label: '参与人', field: 'player' },
        { name: 'language', align: 'center', label: '项目语言', field: 'language' },
        { name: 'node', align: 'center', label: '项目节点', field: 'node' },
        { name: 'status', align: 'center', label: '状态', field: 'status' },
        { name: 'actions', align: 'center', label: '操作', field: 'actions' },
    ]
})
const {
    pagination,
    queryParams,
    pageOptions,
    GqaDictShow,
    GqaAvatar,
    loading,
    tableData,
    recordDetailDialog,
    showAddForm,
    showEditForm,
    onRequest,
    getTableData,
    handleSearch,
    resetSearch,
    handleFinish,
    handleDelete,
} = useTableData(url)

onMounted(() => {
    getTableData()
})

const projectDetailDialog = ref(null)
const showProjectDetail = (row) => {
    projectDetailDialog.value.show(row)
}
</script>
