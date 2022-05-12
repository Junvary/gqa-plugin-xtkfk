<template>
    <div class="gqa-weapon-node" id="gqa-weapon-node">
        <div ref="nodechart" style="width: 100%; height: 500px"></div>
    </div>
</template>

<script setup>
import useTableData from 'src/composables/useTableData'
import { computed, markRaw, onMounted, ref } from 'vue';
const echarts = require('echarts')
const chart = ref(null)
const nodechart = ref(null)

const url = {
    list: 'public/plugin-xtkfk/get-weapon-node-list',
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
    handleSearch,
    resetSearch,
    handleFinish,
    handleDelete,
} = useTableData(url)

onMounted(() => {
    pagination.value.sortBy = 'created_at'
    onRequest({
        pagination: pagination.value,
        queryParams: queryParams.value
    })
    init()
})

const init = () => {
    let ct = nodechart.value;
    echarts.dispose(ct);
    chart.value = markRaw(echarts.init(ct));
    chart.value.setOption(options.value);
}

const options = computed(() => {
    return {
        title: {
            text: '项目节点',
            subtext: tableData.value && tableData.value.length ? '项目当前节点' : '纯属虚构',
            left: 'center',
        },
        tooltip: {
            trigger: 'item',
        },
        series: [
            {
                name: '节点',
                type: 'pie',
                radius: '70%',
                data:
                    tableData.value && tableData.value.length
                        ? tableData.value
                        : [
                            { value: 200, name: '交付上线' },
                            { value: 1, name: '产品下线' },
                            { value: 30, name: '移交运维' },
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
.gqa-weapon-node {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 100px 0;
}
</style>