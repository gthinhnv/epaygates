const DEFAULT_PAGE_NUMBER = 1;
const DEFAULT_PAGE_SIZE = 10;

var $table = $('#static-page-list');
var $delete = $('#delete');
var $applyFilter = $('#apply-filter');
var selections = [];

let statuses = [];
let pageNumber = DEFAULT_PAGE_NUMBER;
let pageSize = DEFAULT_PAGE_SIZE;

function responseHandler(res) {
    $.each(res.data, function (i, row) {
        row.state = $.inArray(row.id, selections) !== -1
    });
    return {
        rows: res.data,
        total: res.total
    };
}

function operateFormatter(value, row, index) {
    return [
        `<a class="table-action-item update" href="/staticPages/${row.id}/update" title="Update">
            <i class="fa fa-edit"></i>
        </a>`
    ].join('');
}

window.operateEvents = {
    'click .delete': function (e, value, row, index) {
        Swal.fire({
            title: 'Are you sure?',
            text: "You won't be able to revert this!",
            icon: 'warning',
            showCancelButton: true,
            confirmButtonText: 'Yes, delete it!'
        }).then((result) => {
            if (result.isConfirmed) {
                deletePages([row.id]);
            }
        });
    }
}

function initTable() {
    const urlParams = getUrlParams();
    const limit = parseInt(urlParams['limit']);
    const offset = parseInt(urlParams['offset']);

    if (limit > 0) {
        pageNumber = Math.floor(offset / limit) + 1;
    }

    if (limit > 0) {
        pageSize = limit;
    }

    $table.bootstrapTable('destroy').bootstrapTable({
        ajaxOptions: function () {
            return {
                authRequired: true,
                xhrFields: { withCredentials: true },
                crossDomain: true
            };
        },
        formatLoadingMessage: function () {
            return '';
        },
        pageNumber: pageNumber,
        pageSize: pageSize,
        fixedColumns: true,
        fixedNumber: 1,
        fixedRightNumber: 1,
        columns: [
            {
                field: 'state',
                checkbox: true,
                align: 'center',
                valign: 'middle'
            }, {
                field: 'id',
                title: 'ID',
                sortable: true,
                align: 'center',
                valign: 'middle'
            }, {
                field: 'title',
                title: trans.TITLE || 'Title',
            }, {
                field: 'slug',
                title: trans.SLUG || 'Slug',
            }, {
                field: 'sortOrder',
                title: trans.SORT_ORDER || 'Sort Order',
            }, {
                field: 'pageType',
                title: trans.PAGE_TYPE || 'pageType',
                formatter: function (value, row, index) {
                    let html = `<select class="table-field-page-type" disabled>`;
                    if (row.pageTypeItem && row.pageTypeItem.id >= 0) {
                        html += `<option value="${row.pageTypeItem.id}" selected>${row.pageTypeItem.text}</option>`;
                    }
                    html += '</select>';
                    return html;
                }
            }, {
                field: 'status',
                title: trans.STATUS || 'Status',
                formatter: function (value, row, index) {
                    let html = `<select class="table-field-status" onchange="updateStatus(${row.id}, $(this).val())">`;
                    if (row.statusItem && row.statusItem.id >= 0) {
                        html += `<option value="${row.statusItem.id}" selected>${row.statusItem.text}</option>`;
                    }
                    html += '</select>';
                    return html;
                }
            }, {
                field: 'createdAt',
                title: trans.CREATED_UPDATED_AT || 'Created/Updated At',
                sortable: false,
                editable: false,
                valign: 'middle',
                align: 'center',
                formatter: function (value, row, index) {
                    return `${moment(row.createdAt).format('DD-MM-YYYY HH:mm:ss')}<br/>${moment(row.updatedAt).format('DD-MM-YYYY HH:mm:ss')}`;
                }
            }, {
                field: 'createdBy',
                title: trans.CREATED_UPDATED_BY || 'Created/Updated By',
                sortable: false,
                editable: false,
                valign: 'middle',
                align: 'center',
                formatter: function (value, row, index) {
                    return `${(row.createdByUser || {}).userName}<br/>${(row.updatedByUser || {}).userName}`;
                }
            }, {
                field: 'operate',
                title: trans.ACTIONS || 'Actions',
                align: 'center',
                clickToSelect: false,
                events: window.operateEvents,
                formatter: operateFormatter
            }
        ]
    });
    $table.on('post-body.bs.table', function (data) {
        selections = getIdSelections();
        $('#static-page-list .table-field-status').select2({
            width: '180px',
            allowClear: false,
            minimumResultsForSearch: Infinity,
            placeholder: 'Select status',
            dropdownParent: $('#static-page-list'),
            ajax: {
                url: `${config.apiAddress}/v1/common/statuses`,
                headers: {
                    'Authorization': `Bearer ${getAccessToken()}`,
                },
                delay: 300,
                data: function (params) {
                    var query = {
                        name: params.term,
                    }
    
                    return query;
                },
                processResults: function (res) {
                    return {
                        results: res.data
                    };
                }
            },
            escapeMarkup: function (markup) {
                return markup;
            },
            templateResult: function (data) {
                return data.text;
            },
            templateSelection: function (data) {
                return data.text;
            }
        });

        $('#static-page-list .table-field-page-type').select2({
            width: '180px',
            allowClear: false,
            minimumResultsForSearch: Infinity,
            placeholder: 'Select page type',
            dropdownParent: $('#static-page-list'),
            ajax: {
                url: `${config.apiAddress}/v1/common/pageTypes`,
                headers: {
                    'Authorization': `Bearer ${getAccessToken()}`,
                },
                delay: 300,
                data: function (params) {
                    var query = {
                        name: params.term,
                    }
    
                    return query;
                },
                processResults: function (res) {
                    return {
                        results: res.data
                    };
                }
            },
            escapeMarkup: function (markup) {
                return markup;
            },
            templateResult: function (data) {
                return data.text;
            },
            templateSelection: function (data) {
                return data.text;
            }
        });
    });
    $table.on('check.bs.table uncheck.bs.table ' + 'check-all.bs.table uncheck-all.bs.table', function () {
        if (!$table.bootstrapTable('getSelections').length) {
            $delete.addClass('disabled');
            $delete.prop('disabled', true);
        } else {
            $delete.removeClass('disabled');
            $delete.prop('disabled', false);
        }

        // save your data, here just save the current page
        selections = getIdSelections();
        // push or splice the selections if you want to save all data selections
    });
    $table.on('all.bs.table', function (e, name, args) {
        // console.log(name, args)
    });
    $delete.click(function () {
        Swal.fire({
            title: 'Are you sure?',
            text: "You won't be able to revert this!",
            icon: 'warning',
            showCancelButton: true,
            confirmButtonText: 'Yes, delete it!'
        }).then((result) => {
            if (result.isConfirmed) {
                var ids = getIdSelections();
                deletePages(ids);
            }
        });
    });
}

function updateStatus(id, status) {
    const endPoint = `${config.apiAddress}/v1/staticPages/${id}/update`;
    $.ajax({
        type: 'POST',
        url: endPoint,
        authRequired: true,          // 👈 marks this request for interception
        xhrFields: { withCredentials: true },
        crossDomain: true,
        data: JSON.stringify({ status: parseInt(status, 10), fields: ['status'] }),
        success: function (data) {
            toastr['success'](`Updated page status successfully!`);
        },
        error: function (xhr, status, error) {
            const message = (xhr.responseJSON || {}).message || `Failed to updated page status. Please contact admin.`;
            toastr['error'](message);
        },
        contentType: 'application/json',
        dataType: 'json'
    });
}

function deletePages(ids) {
    const endPoint = `${config.apiAddress}/v1/staticPages/delete`;
    $.ajax({
        type: 'DELETE',
        url: endPoint,
        authRequired: true,          // 👈 marks this request for interception
        xhrFields: { withCredentials: true },
        crossDomain: true,
        data: JSON.stringify({ ids: ids }),
        success: function (data) {
            toastr['success'](`Deleted admin page with ids = ${ids} successfully!`);
            $table.bootstrapTable('remove', {
                field: 'id',
                values: ids
            });
            if (getIdSelections().length == 0) {
                $delete.addClass('disabled');
                $delete.prop('disabled', true);
            }
        },
        error: function (xhr, status, error) {
            const message = (xhr.responseJSON || {}).message || `Failed to delete page with ids = ${ids}. Please contact admin.`;
            toastr['error'](message);
        },
        contentType: 'application/json',
        dataType: 'json'
    });
}

$(function () {
    initTable();

    $applyFilter.click(function () {
        const urlParams = getUrlParams();
        const tableOptions = $table.bootstrapTable('getOptions');
        var sortName = tableOptions.sortName;
        var sortOrder = tableOptions.sortOrder;
        const limit = parseInt(urlParams['limit']);
        const offset = parseInt(urlParams['offset']);

        if (urlParams['sort']) {
            const sortArr = urlParams['sort'].split('.');
            if (sortArr.length == 2) {
                sortName = sortArr[0];
                sortOrder = sortArr[1];
            }
        }

        if (limit > 0) {
            pageNumber = Math.floor(offset / limit) + 1;
        }
        $table.bootstrapTable('refreshOptions', {
            pageNumber: pageNumber,
            sortName: sortName,
            sortOrder: sortOrder
        });
    });

    $('.reset-filter-onchange').change(function () {
        pageNumber = 1;
        setToFirstPage();
    });
})