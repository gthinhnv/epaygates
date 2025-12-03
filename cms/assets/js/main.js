/**
 * Common
 */
function getAccessToken() {
    return getCookie(config.accessTokenKey);
}

function ensureAccessToken() {
    const accessToken = getAccessToken();
    const isRefreshTokenAvailable = getCookie(`${config.refreshTokenKey}_available`) == '1';

    // If token exists
    if (accessToken && accessToken.length > 0) {
        return Promise.resolve(accessToken);
    }

    if (!isRefreshTokenAvailable) {
        return Promise.reject(new Error('No refresh token available'));
    }

    const endPoint = '/auth/refreshToken';

    return new Promise((resolve, reject) => {
        $.ajax({
            type: 'POST',
            url: endPoint,
            xhrFields: {
                withCredentials: true   // 👈 required for cookies
            },
            crossDomain: true,
            contentType: 'application/json',
            dataType: 'json',
            success: function (res) {
                console.log('refresh token result', res);
                if (!res || res.code !== 0 || !res.data?.accessToken) {
                    return reject(new Error('Failed to refresh token'));
                }

                resolve(res.data.accessToken);
            },
            error: function (xhr, status, err) {
                console.error('refresh token failed', err);
                reject(new Error('Failed to refresh token'));
            }
        });
    });
}

function getParameterByName(name, url = window.location.href) {
    name = name.replace(/[\[\]]/g, '\\$&');
    var regex = new RegExp('[?&]' + name + '(=([^&#]*)|&|#|$)'),
        results = regex.exec(url);
    if (!results) return null;
    if (!results[2]) return '';
    return decodeURIComponent(results[2].replace(/\+/g, ' '));
}

function getIdSelections() {
    return $.map($table.bootstrapTable('getSelections'), function (row) {
        return row.id;
    });
}

function getIds() {
    return $.map($table.bootstrapTable('getData'), function (row) {
        return row.id;
    });
}

function getUrlParams() {
    const search = location.search.substring(1);
    const urlParams = search ? JSON.parse('{"' + search.replace(/&/g, '","').replace(/=/g, '":"') + '"}', function (key, value) { return key === "" ? value : decodeURIComponent(value) }) : {}
    return urlParams;
}

function initTableFilters() {
    const urlParams = getUrlParams();

    $('#filter').find('input[name], select[name]').each(function () {
        let v = urlParams[$(this).attr('name')];
        if (v == undefined || v == null) {
            v = $(this).attr('default') || '';
        }
        if ($(this).hasClass('is-select2')) {
            if ($(this).children().length === 0) {
                $(this).append(`<option value="${v}">${v}</option>`);
            }
            $(this).val(v).trigger('change');
        } else if ($(this).hasClass('is-datetimepicker')) {
            const dateTimeFormat = $(this).attr('format') || 'DD-MM-YYYY HH:mm';
            const time = $(this).attr('time');
            const d = moment(v);
            if (time && time.length > 0) {
                const timeArr = time.split(':');
                if (timeArr.length > 0) {
                    d.set("hour", timeArr[0])
                }
                if (timeArr.length > 1) {
                    d.set("minute", timeArr[1])
                }
                if (timeArr.length > 2) {
                    d.set("second", timeArr[2])
                }
            }
            $(this).val(d.isValid() ? d.format(dateTimeFormat) : '');
        } else {
            $(this).val(v);
        }
    });
}

function queryParams(params) {
    $('#filter').find('input[name], select[name]').each(function () {
        var v = $(this).val();
        if ($(this).hasClass('is-datetimepicker')) {
            const dateTimeFormat = $(this).attr('format') || 'DD-MM-YYYY HH:mm';
            const time = $(this).attr('time');
            const d = moment(v, dateTimeFormat);
            if (time && time.length > 0) {
                const timeArr = time.split(':');
                if (timeArr.length > 0) {
                    d.set("hour", timeArr[0])
                }
                if (timeArr.length > 1) {
                    d.set("minute", timeArr[1])
                }
                if (timeArr.length > 2) {
                    d.set("second", timeArr[2])
                }
            }
            v = d.isValid() ? d.utc().format('YYYY-MM-DDTHH:mm:ssZ') : '';
        }
        if ($.isArray(v)) {
            v = v.join(',');
        }
        params[$(this).attr('name')] = v;
    });
    // Custom sort
    if (params['sort'] && params['sort'] !== '' && params['order'] && params['order'] !== '') {
        params['sort'] = `${params['sort']}.${params['order']}`;
        delete params['order'];
    }
    $.each(params, function (key, value) {
        if (value === undefined || value === null || value === '') {
            delete params[key];
        }
    });
    const newUrl = window.location.protocol + "//" + window.location.host + window.location.pathname + '?' + $.param(params);
    window.history.pushState({ path: newUrl }, '', newUrl);
    return params;
}

function setToFirstPage() {
    setQueryParam('offset', 0);
}

function setQueryParam(key, value) {
    const urlParams = getUrlParams();
    urlParams[key] = value;
    const newUrl = window.location.protocol + "//" + window.location.host + window.location.pathname + '?' + $.param(urlParams);
    window.history.pushState({ path: newUrl }, '', newUrl);
}

function setFormValues(selector, data = {}) {
    $(selector).find('input[name],select[name],textarea[name]').each(function () {
        let v = data[$(this).attr('name')];
        if (v == undefined || v == null) {
            v = $(this).attr('default') || '';
        }
        if ($(this).hasClass('is-select2')) {
            $(this).val(v).trigger('change');
        } else if ($(this).hasClass('html-editor')) {
            CKEDITOR.instances[$(this).attr('name')].setData(v);
        } else if ($(this).hasClass('is-icheck')) {
            $(this).iCheck(v ? 'check' : 'uncheck');
        } else if ($(this).hasClass('is-datetimepicker')) {
            const d = moment(v);
            $(this).val(d.isValid() ? d.format('DD-MM-YYYY HH:mm') : '');
        } else {
            $(this).val(v);
        }
    });
}

function addFieldError($input, errMsg) {
    const $parent = $input.closest('.form-group');
    $input.addClass('form-control-danger');
    $input.attr('has-error', '1');
    if ($input.hasClass('is-select2')) {
        $input.change(function () {
            clearFieldError($(this));
        });
    } else if ($input.hasClass('html-editor')) {
        CKEDITOR.instances[$input.attr('name')].document.on('keydown', function (event) {
            clearFieldError($input);
        });
    } else {
        $input.keydown(function () {
            clearFieldError($(this));
        });
    }
    $parent.find('.err-msg').remove();
    $parent.addClass('has-danger');
    $parent.append(`<div class="col-form-label err-msg">${errMsg}</div>`);
}

function clearFieldError($input) {
    if ($input.attr('has-error') == 1) {
        const $parent = $input.closest('.form-group');
        $input.removeClass('form-control-danger');
        $input.removeAttr('has-error');
        $parent.removeClass('has-danger');
        $parent.find('.err-msg').remove();
    }
}

function setCookie(cname, cvalue, seconds) {
    const d = new Date();
    d.setTime(d.getTime() + (seconds * 1000));
    let expires = 'expires=' + d.toUTCString();
    document.cookie = cname + '=' + cvalue + ';' + expires + ';path=/';
}

function getCookie(cname) {
    let name = cname + '=';
    let ca = document.cookie.split(';');
    for (let i = 0; i < ca.length; i++) {
        let c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return '';
}

function toSlug(source, des) {
    const title = source.value;
    if (!title) {
        return false;
    }
    const slug = slugify(title, {
        lower: true,      // convert to lower case, defaults to `false`
        strict: true     // strip special characters except replacement, defaults to `false`
    });
    document.querySelector(des).value = slug.toLowerCase();
}

function getJwtTokenRemainingTime(token) {
    const tokenParts = token.split('.');
    if (tokenParts.length !== 3) {
        return -1; // Invalid token format
    }

    const payload = JSON.parse(atob(tokenParts[1]));
    if (!payload.hasOwnProperty('exp')) {
        return -1; // Token doesn't have expiration claim
    }

    const expirationTime = payload.exp * 1000; // Convert to milliseconds
    const currentTime = new Date().getTime();

    return expirationTime - currentTime;
}

$.ajaxPrefilter(function(options, originalOptions, jqXHR) {
    if (options.authRequired) {
        // Abort the original request so only the retried one fires
        jqXHR.abort();

        var dfd = $.Deferred();

        ensureAccessToken().then(function (token) {
            // clone options safely to avoid mutating original
            var newOptions = $.extend(true, {}, options, {
                headers: $.extend({}, options.headers, {
                    Authorization: "Bearer " + token
                }),
                authRequired: false // prevent infinite loop
            });

            // retry with updated headers
            $.ajax(newOptions).then(dfd.resolve, dfd.reject);

        }).catch(function (err) {
            console.warn("Token refresh failed, calling API without token", err);

            // retry original request without auth
            var fallbackOptions = $.extend(true, {}, options, {
                authRequired: false // prevent infinite loop
            });

            $.ajax(fallbackOptions).then(dfd.resolve, dfd.reject);
        });

        // IMPORTANT: return the deferred's promise
        // so the calling code only sees the retried request,
        // not the aborted one
        return dfd.promise(jqXHR);
    }
});

$(document).ready(function () {
    Waves.init();
    Waves.attach('.flat-buttons', ['waves-button']);
    Waves.attach('.float-buttons', ['waves-button', 'waves-float']);
    Waves.attach('.float-button-light', ['waves-button', 'waves-float', 'waves-light']);
    Waves.attach('.flat-buttons', ['waves-button', 'waves-float', 'waves-light', 'flat-buttons']);
    $('.theme-loader').animate({ opacity: '0' }, 1200);
    setTimeout(function () {
        $('.theme-loader').remove();
    }, 2000);

    // setInterval(function() {
    //     checkAndRefreshToken();
    // }, 300000); // 5mins
});
$(document).on('select2:open', () => {
    document.querySelector('.select2-search__field').focus();
});