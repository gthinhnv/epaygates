$(function () {
    $('#static-page-form.form-create').submit(function (event) {
        event.preventDefault();
        const formData = getFormData($(this));
    
        $.ajax({
            type: 'POST',
            url: `${config.apiAddress}/v1/staticPages/create`,
            authRequired: true,          // 👈 marks this request for interception
            xhrFields: { withCredentials: true },
            crossDomain: true,
            data: JSON.stringify(formData),
            success: function (data) {
                toastr['success'](`A new page was created successfully!`);
                window.location.replace('/pages');
            },
            error: function (xhr, status, error) {
                const resJson = xhr.responseJSON || {};
                const errs = resJson.errs || {};
                $('#page-form').find('input[name], select[name], textarea[name]').each(function () {
                    const name = $(this).attr('name');
                    const errMsg = errs[name];
                    if (errMsg && errMsg != '') {
                        addFieldError($(this), errMsg);
                    }
                });
                toastr['error'](resJson.message || `An error occurred!`);
            },
            contentType: 'application/json',
            dataType: 'json'
        });
    
        return false;
    });
});