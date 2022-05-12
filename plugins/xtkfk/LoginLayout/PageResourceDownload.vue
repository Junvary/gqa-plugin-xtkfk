<template>
    <div class="q-pa-md row items-center justify-center gqa-resource-download" id="gqa-resource-download">
        <q-table row-key="id" separator="cell" :rows="dataList" :columns="columns" v-model:pagination="pagination"
            :rows-per-page-options="pageOptions" :loading="loading" @request="onRequest"
            style="min-width: 40vw; max-width: 45vw">

            <template v-slot:top>
                <span class="text-h6">
                    资源查阅
                </span>
                <q-space />
                <!-- <q-input dense v-model="queryParams.title" label="标题" />
                <q-btn dense color="primary" @click="handleSearch" label="搜索" style="margin: 0 10px" /> -->
                <q-btn dense round push icon="cached" @click="onRequest" />
            </template>

            <template v-slot:body-cell-title="props">
                <q-td :props="props">
                    <q-btn flat dense rounded glossy @click="showDetail(props.row)">
                        {{ props.row.title }}
                    </q-btn>
                </q-td>
            </template>

            <template v-slot:body-cell-created_at="props">
                <q-td :props="props">
                    {{ showDateTime(props.row.created_at) }}
                </q-td>
            </template>

            <template v-slot:body-cell-created_by="props">
                <q-td :props="props">
                    <GqaShowName v-if="props.row.created_by_user" :customNameObject="props.row.created_by_user" />
                </q-td>
            </template>

        </q-table>

        <q-dialog v-model="showDetailVisible">
            <q-card style="width: 1400px; max-width: 80vw;">
                <q-card-section class="row items-center q-pb-none">
                    <div class="text-h6">
                        {{ detail.title }}
                    </div>
                    <q-space />
                    <q-btn icon="close" flat round dense v-close-popup />
                </q-card-section>

                <q-separator />

                <q-card-section v-html="detail.content" style="max-height: 80vh" class="scroll" />

                <q-separator />

                <q-list bordered separator v-if="detail.attachment">
                    <q-item-label header class="row items-center justify-between">
                        <span>附件列表:</span>
                    </q-item-label>
                    <q-separator />
                    <q-item clickable v-for="(item, index) in JSON.parse(detail.attachment)" :key="index"
                        @click="handleDownload(item)">
                        <q-item-section>
                            {{ item.filename }}
                        </q-item-section>
                    </q-item>
                </q-list>
            </q-card>
        </q-dialog>
    </div>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { useQuasar } from 'quasar'
import { downloadAction } from 'src/api/manage'
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { DictOptions } from 'src/utils/dict'
import { FormatDateTime } from 'src/utils/date'

const $q = useQuasar()
const { t } = useI18n()
const url = {
    list: 'public/plugin-xtkfk/get-download-list',
}
const columns = [
    { name: 'title', align: 'center', label: '资源标题', field: 'title' },
    { name: 'created_by', align: 'center', label: '发布人', field: 'created_by' },
    { name: 'created_at', align: 'center', label: '发布时间', field: 'created_at' },
]
const {
    showDateTime,
    dictOptions,
    pagination,
    queryParams,
    pageOptions,
    GqaDictShow,
    GqaAvatar,
    GqaShowName,
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

const showDetailVisible = ref(false)
const detail = ref({})
onMounted(() => {
    pagination.value.sortBy = 'created_at'
    onRequest({
        pagination: pagination.value,
        queryParams: queryParams.value
    })
})

const dataList = computed(() => {
    if (tableData.value && tableData.value.length) {
        return tableData.value
    } else {
        return [{
            title: 'Gin-Quasar-Admin',
            content: 'Gin-Quasar-Admin',
        },]
    }
})
const showDetail = (row) => {
    detail.value = row
    showDetailVisible.value = true
}
const handleDownload = (item) => {
    downloadAction(item.fileUrl.substring(11), item.filename)
}
</script>

<style lang="scss" scoped>
.gqa-resource-download {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
}
</style>