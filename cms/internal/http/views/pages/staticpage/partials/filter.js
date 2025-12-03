$(function () {
    initTableFilters();

    $('#filter .status-selection').select2({
        width: '100%',
        allowClear: true,
        minimumResultsForSearch: Infinity,
        placeholder: 'Select status',
        dropdownParent: $('#filter'),
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

    const selectedStatus = parseInt($('#filter .status-selection').val());
    if (selectedStatus >= 0) {
        $.ajax({
            type: 'GET',
            url: `${config.apiAddress}/v1/common/statuses`,
            authRequired: true,          // 👈 marks this request for interception
            xhrFields: { withCredentials: true },
            crossDomain: true,
            success: function (res) {
                const data = res && res.data || [];
                const selectedItem = data.find(x => x.id === selectedStatus);
                if (selectedItem) {
                    $('#filter .status-selection').empty().append(`<option value="${selectedItem.id}" selected>${selectedItem.text}</option>`).trigger('change');
                }
            },
            error: function (xhr, status, error) {
                const message = (xhr.responseJSON || {}).message || `Failed to get page status. Please contact admin.`;
                toastr['error'](message);
            },
            contentType: 'application/json',
            dataType: 'json'
        });
    } else {
        $('#filter .status-selection').empty().trigger('change');
    }
});