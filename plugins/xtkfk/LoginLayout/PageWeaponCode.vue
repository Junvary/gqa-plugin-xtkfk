<template>
    <div class="gqa-weapon-code" id="gqa-weapon-code">
        <div ref="codechart" style="width: 100%; height: 500px"></div>
    </div>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { computed, markRaw, onMounted, ref, watch } from 'vue';
const echarts = require('echarts')
const chart = ref(null)
const codechart = ref(null)

const url = {
    list: 'public/plugin-xtkfk/get-weapon-language-list',
}
const {
    showDateTime,
    dictOptions,
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

onMounted(async () => {
    pagination.value.sortBy = 'created_at'
    await getTableData()
    init()
})

const init = () => {
    let ct = codechart.value;
    echarts.dispose(ct);
    chart.value = markRaw(echarts.init(ct));
    chart.value.setOption(options.value);
}

const options = computed(() => {
    return {
        title: {
            text: '项目语言',
            subtext: tableData.value && tableData.value.length ? '项目使用语言' : '纯属虚构',
            left: 'center',
        },
        tooltip: {
            trigger: 'item',
        },
        series: [
            {
                name: '语言',
                type: 'pie',
                radius: '70%',
                data:
                    tableData.value && tableData.value.length
                        ? tableData.value
                        : [
                            { value: 1048, name: 'Java' },
                            { value: 735, name: 'C#' },
                            { value: 580, name: 'C++' },
                            { value: 484, name: 'Python' },
                            { value: 300, name: 'Golang' },
                        ],
                emphasis: {
                    itemStyle: {
                        shadowBlur: 10,
                        shadowOffsetX: 0,
                        shadowColor: 'rgba(0, 0, 0, 0.5)',
                    },
                },
                label: {
                    show: true,
                    // formatter: '{b}:{c}',
                },
            },
        ],
    }
})
</script>

<style lang="scss" scoped>
.gqa-weapon-code {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 100px 0;
}
</style>