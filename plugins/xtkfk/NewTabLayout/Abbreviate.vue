<template>
    <div class="q-pa-md">
        <div class="text-center text-h4 text-bold">
            简称短语汇总
        </div>
        <q-stepper flat v-model="step" ref="stepper" color="primary" header-nav animated>
            <q-step v-for="(item, index) in dataList" :name="index" :title="'大类：' + item.top" icon="star">
                <q-timeline color="primary">
                    <q-timeline-entry v-for="a in item.children" :title="a.title + ' ' + a.from">
                        <div v-html="a.content"></div>
                    </q-timeline-entry>
                </q-timeline>
            </q-step>
        </q-stepper>
    </div>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { computed, onMounted, ref } from 'vue';

const step = ref(0)
const url = {
    list: 'public/plugin-xtkfk/get-abbreviate-list',
}
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
    getTableData,
    handleSearch,
    resetSearch,
    handleFinish,
    handleDelete,
} = useTableData(url)

onMounted(() => {
    pagination.value.rowsPerPage = 9999999
    pagination.value.sortBy = 'created_at'
    pagination.value.descending = true
    getTableData()
})
const dataList = computed(() => {
    const topList = []
    for (let i of tableData.value) {
        if (topList.filter(item => item.top === i.top).length === 0) {
            topList.push({
                top: i.top,
                children: [i]
            })
        } else {
            topList.filter(item => item.top === i.top)[0].children.push(i)
        }
    }
    topList.sort(function (a, b) {
        return a.top - b.top
    })
    return topList
})
</script>