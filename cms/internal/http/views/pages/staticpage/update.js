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
                window.location.replace('/staticPages');
            },
            error: function (xhr, status, error) {
                const resJson = xhr.responseJSON || {};
                const errors = resJson.data || {};
                errors.forEach(e => {
                    const { field, message } = e;

                    // Escape special characters for jQuery attribute selector
                    const safeField = field.replace(/([!"#$%&'()*+,./:;<=>?@[\\\]^`{|}~])/g, "\\$1");

                    const $input = $('#static-page-form').find(`[name="${safeField}"]`).first();

                    if ($input.length) {
                        addFieldError($input, message);
                    } else {
                        console.warn(`Field not found: ${field}`);
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