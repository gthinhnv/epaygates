function getFormData($form) {
    var formData = $form.serializeArray()
        .reduce(function (json, { name, value }) {
            json[name] = value;
            return json;
        }, {});

    return {
        title: formData.title || '',
        slug: formData.slug || '',
        sortOrder: parseInt(formData.sortOrder) || 0,
        content: CKEDITOR.instances.content.getData(),
        pageType: parseInt(formData.pageType) || 0,
        seo: {
            metaTitle: formData.metaTitle || '',
            metaDescription: formData.metaDescription || '',
            metaKeywords: formData.metaKeywords || '',
            schema: formData.schema || '',
        },
        adsPlatform: parseInt(formData.adsPlatform) || 0,
        status: parseInt(formData.status) || 0
    }
}

$(function () {
    CKEDITOR.replace('content', {
        toolbarGroups: [
            { name: 'document', groups: ['mode', 'document', 'doctools'] },
            { name: 'clipboard', groups: ['clipboard', 'undo'] },
            { name: 'editing', groups: ['find', 'selection', 'spellchecker', 'editing'] },
            { name: 'forms', groups: ['forms'] },
            '/',
            { name: 'basicstyles', groups: ['basicstyles', 'cleanup'] },
            { name: 'paragraph', groups: ['list', 'indent', 'blocks', 'align', 'bidi', 'paragraph'] },
            { name: 'links', groups: ['links'] },
            { name: 'insert', groups: ['insert'] },
            '/',
            { name: 'styles', groups: ['styles'] },
            { name: 'colors', groups: ['colors'] },
            { name: 'tools', groups: ['tools'] },
            { name: 'others', groups: ['others'] },
            { name: 'about', groups: ['about'] }
        ],
        removeButtons: 'Save,NewPage,Templates,PasteText,PasteFromWord,ImageButton,Button,Select,Textarea,TextField,Form,Checkbox,Radio,HiddenField,CopyFormatting,RemoveFormat,Subscript,Language,BidiRtl,BidiLtr,Flash,PageBreak,Iframe,Font,ShowBlocks,About',
        filebrowserBrowseUrl: '/media/file-manager?sf=/content/page',
        filebrowserUploadUrl: '/media/editor-upload-file?sf=/content/page',
        height: '500px',
        on: {
            instanceReady: function () {
                const dtd = CKEDITOR.dtd;
    
                for ( var e in CKEDITOR.tools.extend( {}, 
                    dtd.$block, 
                    dtd.$listItem, 
                    dtd.$tableContent ) 
                ) {
                    this.dataProcessor.writer.setRules( e, {
                        indent: false,
                        breakBeforeOpen: false,
                        breakAfterOpen: false,
                        breakBeforeClose: false,
                        breakAfterClose: false
                    });
                }
            }
        }
    });

    $('#static-page-form .status-selection').select2({
        width: '100%',
        allowClear: false,
        minimumResultsForSearch: Infinity,
        placeholder: 'Select status',
        dropdownParent: $('#static-page-form'),
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

    $('#static-page-form .page-type-selection').select2({
        width: '100%',
        allowClear: false,
        minimumResultsForSearch: Infinity,
        placeholder: 'Select page type',
        dropdownParent: $('#static-page-form'),
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

    $('#static-page-form .ads-platform-selection').select2({
        width: '100%',
        allowClear: false,
        minimumResultsForSearch: Infinity,
        placeholder: 'Select Ads Platform',
        dropdownParent: $('#static-page-form'),
        ajax: {
            url: `${config.apiAddress}/v1/common/adsPlatforms`,
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