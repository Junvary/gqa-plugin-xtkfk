<template>
    <q-page padding>
        <div class="row q-gutter-md items-center" style="margin-bottom: 10px">
            <q-input style="width: 20%" v-model="queryParams.title" label="标题" />
            <q-btn color="primary" @click="handleSearch" label="搜索" />
            <q-btn color="primary" @click="resetSearch" label="重置" />
        </div>

        <q-table row-key="id" separator="cell" :rows="tableData" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest">

            <template v-slot:top="props">
                <q-btn color="primary" @click="showAddForm()" label="新增荣誉认证" />
                <q-space />
                <q-btn flat round dense :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
                    @click="props.toggleFullscreen" class="q-ml-md" />
            </template>

            <template v-slot:body-cell-status="props">
                <q-td :props="props">
                    <GqaDictShow dictName="statusOnOff" :dictCode="props.row.status" />
                </q-td>
            </template>

            <template v-slot:body-cell-created_by="props">
                <q-td :props="props">
                    <GqaShowName v-if="props.row.created_by_user" :customNameObject="props.row.created_by_user" />
                </q-td>
            </template>

            <template v-slot:body-cell-created_at="props">
                <q-td :props="props">
                    {{ showDateTime(props.row.created_at) }}
                </q-td>
            </template>

            <template v-slot:body-cell-actions="props">
                <q-td :props="props">
                    <div class="q-gutter-xs">
                        <q-btn color="primary" @click="showEditForm(props.row)" label="编辑" />
                        <q-btn color="negative" @click="handleDelete(props.row)" label="删除" />
                    </div>
                </q-td>
            </template>
        </q-table>
        <recordDetail ref="recordDetailDialog" @handleFinish="handleFinish" />
    </q-page>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { useQuasar } from 'quasar'
import { postAction } from 'src/api/manage'
import { computed, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import recordDetail from './modules/recordDetail'
import GqaShowName from 'src/components/GqaShowName'

const $q = useQuasar()
const { t } = useI18n()
const url = {
    list: 'plugin-xtkfk/get-honour-list',
    delete: 'plugin-xtkfk/delete-honour',
}
const columns = computed(() => {
    return [
        { name: 'title', align: 'center', label: '荣誉认证标题', field: 'title' },
        { name: 'created_by', align: 'center', label: '发布人', field: 'created_by' },
        { name: 'created_at', align: 'center', label: '发布时间', field: 'created_at' },
        { name: 'status', align: 'center', label: '状态', field: 'status' },
        { name: 'actions', align: 'center', label: '操作', field: 'actions' },
    ]
})
const {
    showDateTime,
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
    handleSearch,
    resetSearch,
    handleFinish,
    handleDelete,
} = useTableData(url)

onMounted(() => {
    onRequest({
        pagination: pagination.value,
        queryParams: queryParams.value
    })
})
</script>
