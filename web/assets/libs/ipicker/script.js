class IPicker {
	constructor(config) {
		this.config = config;

		this.initHTML();

		this.createExistImgsPreview();

		if (!window.pixie) {
			// default-settings.ts
			window.pixie = new Pixie({
				googleFontsApiKey: 'AIzaSyC0oTO0Ni61rDRSfI5OTO-_SzeoToWnUH0',
				toolbar: {
					hideOpenButton: true,
					hideCloseButton: false,
				},
				ui: {
					theme: 'light',
					mode: 'overlay',
					visible: false,
					nav: {
						replaceDefault: false,
						position: 'bottom'
					},
					openImageDialog: {
						show: false
					}
				},
				textureSize: 8000,
				baseUrl: '/assets/libs/ipicker/pixie-2.1.3/pixie/',
				tools: {
					crop: {
						replaceDefault: true,
						items: ['1.91:1', '16:9', '3:2', '5:4', "1:1", '4:5', '2:3', '9:16', '1:2.1']
					}
				},
				onLoad: function () {
					console.log('Pixie is ready');
					let globalLoadingHtml = `<div class="global-spinner">
						<style>.global-spinner {display: none; align-items: center; justify-content: center; z-index: 9999; background: rgba(0, 0, 0, 0.28); position: absolute; top: 0; left: 0; width: 100%; height: 100%;}</style>
						<style>.la-ball-spin-clockwise,.la-ball-spin-clockwise>div{position:relative;-webkit-box-sizing:border-box;-moz-box-sizing:border-box;box-sizing:border-box}.la-ball-spin-clockwise{display:block;font-size:0;color:#ffffff}.la-ball-spin-clockwise.la-dark{color:#333}.la-ball-spin-clockwise>div{display:inline-block;float:none;background-color:currentColor;border:0 solid currentColor}.la-ball-spin-clockwise{width:32px;height:32px}.la-ball-spin-clockwise>div{position:absolute;top:50%;left:50%;width:8px;height:8px;margin-top:-4px;margin-left:-4px;border-radius:100%;-webkit-animation:ball-spin-clockwise 1s infinite ease-in-out;-moz-animation:ball-spin-clockwise 1s infinite ease-in-out;-o-animation:ball-spin-clockwise 1s infinite ease-in-out;animation:ball-spin-clockwise 1s infinite ease-in-out}.la-ball-spin-clockwise>div:nth-child(1){top:5%;left:50%;-webkit-animation-delay:-.875s;-moz-animation-delay:-.875s;-o-animation-delay:-.875s;animation-delay:-.875s}.la-ball-spin-clockwise>div:nth-child(2){top:18.1801948466%;left:81.8198051534%;-webkit-animation-delay:-.75s;-moz-animation-delay:-.75s;-o-animation-delay:-.75s;animation-delay:-.75s}.la-ball-spin-clockwise>div:nth-child(3){top:50%;left:95%;-webkit-animation-delay:-.625s;-moz-animation-delay:-.625s;-o-animation-delay:-.625s;animation-delay:-.625s}.la-ball-spin-clockwise>div:nth-child(4){top:81.8198051534%;left:81.8198051534%;-webkit-animation-delay:-.5s;-moz-animation-delay:-.5s;-o-animation-delay:-.5s;animation-delay:-.5s}.la-ball-spin-clockwise>div:nth-child(5){top:94.9999999966%;left:50.0000000005%;-webkit-animation-delay:-.375s;-moz-animation-delay:-.375s;-o-animation-delay:-.375s;animation-delay:-.375s}.la-ball-spin-clockwise>div:nth-child(6){top:81.8198046966%;left:18.1801949248%;-webkit-animation-delay:-.25s;-moz-animation-delay:-.25s;-o-animation-delay:-.25s;animation-delay:-.25s}.la-ball-spin-clockwise>div:nth-child(7){top:49.9999750815%;left:5.0000051215%;-webkit-animation-delay:-.125s;-moz-animation-delay:-.125s;-o-animation-delay:-.125s;animation-delay:-.125s}.la-ball-spin-clockwise>div:nth-child(8){top:18.179464974%;left:18.1803700518%;-webkit-animation-delay:0s;-moz-animation-delay:0s;-o-animation-delay:0s;animation-delay:0s}.la-ball-spin-clockwise.la-sm{width:16px;height:16px}.la-ball-spin-clockwise.la-sm>div{width:4px;height:4px;margin-top:-2px;margin-left:-2px}.la-ball-spin-clockwise.la-2x{width:64px;height:64px}.la-ball-spin-clockwise.la-2x>div{width:16px;height:16px;margin-top:-8px;margin-left:-8px}.la-ball-spin-clockwise.la-3x{width:96px;height:96px}.la-ball-spin-clockwise.la-3x>div{width:24px;height:24px;margin-top:-12px;margin-left:-12px}@-webkit-keyframes ball-spin-clockwise{0%,100%{opacity:1;-webkit-transform:scale(1);transform:scale(1)}20%{opacity:1}80%{opacity:0;-webkit-transform:scale(0);transform:scale(0)}}@-moz-keyframes ball-spin-clockwise{0%,100%{opacity:1;-moz-transform:scale(1);transform:scale(1)}20%{opacity:1}80%{opacity:0;-moz-transform:scale(0);transform:scale(0)}}@-o-keyframes ball-spin-clockwise{0%,100%{opacity:1;-o-transform:scale(1);transform:scale(1)}20%{opacity:1}80%{opacity:0;-o-transform:scale(0);transform:scale(0)}}@keyframes ball-spin-clockwise{0%,100%{opacity:1;-webkit-transform:scale(1);-moz-transform:scale(1);-o-transform:scale(1);transform:scale(1)}20%{opacity:1}80%{opacity:0;-webkit-transform:scale(0);-moz-transform:scale(0);-o-transform:scale(0);transform:scale(0)}}</style>
						<div class="la-ball-spin-clockwise la-2x">
							<div></div>
							<div></div>
							<div></div>
							<div></div>
							<div></div>
							<div></div>
							<div></div>
							<div></div>
						</div>
					</div>`;
					$(globalLoadingHtml).appendTo('image-editor');
				}.bind(this),
				onSave: function (data, name) {
					this.saveToServer(data);
				}.bind(this)
			});
		}

		this.initEvents();
	}

	initHTML = () => {
		let selectors = document.getElementsByTagName('pixie-editor');
		let ipickerModal = document.getElementById('ipicker-modal');
		if (selectors.length == 0) {
			let selectorHTML = `<pixie-editor></pixie-editor>`;
			$(selectorHTML).appendTo('body');
		}
		if (!ipickerModal) {
			let modalHtml = `<div id="ipicker-modal" class="modal fade" role="dialog" aria-hidden="true">
				<div class="modal-dialog">
					<!-- Modal content-->
					<div class="modal-content">
						<div class="modal-header">
							<h4 class="modal-title">Image Picker</h4>
							<button type="button" class="close" data-dismiss="modal" aria-label="Close">
								<span aria-hidden="true">&times;</span>
							</button>
						</div>
						<div class="modal-body">
							<div class="sources text-center">
								<div class="s-computer">
									<div class="upload-btn-wrapper">
										<div class="file-upload-display">
											<i class="fa fa-cloud-upload-alt" aria-hidden="true"></i>
											<p class="instruction">Drag and drop a file here or click</p>
										</div>
										<input type="file" class="file-upload custom-file-input">
									</div>
								</div>
								<div class="s-library"></div>
							</div>
						</div>
					</div>
				</div>
			</div>`;
			$(modalHtml).appendTo('body');
		}
	}

	initEvents = () => {
		$(this.config.selector + ' .ipicker-btn').click(function () {
			$('#ipicker-modal').modal({
				show: true,
				backdrop: 'static',
				keyboard: false
			});
		});

		$(document).on('change', '.file-upload', function (e) {
			this.readURL(e.target);
		}.bind(this));
	}

	/**
	 * Event on input change
	 */
	async readURL(input) {

		if (input.files && input.files.length) {
			var file = input.files[0];

			if (/^image\/\w+/.test(file.type)) {
				this.imageType = file.type;

				var reader = new FileReader();
				reader.readAsDataURL(file);

				var imgBase64 = await this.imgtoBase64(file);
				window.pixie.resetAndOpenEditor({ image: imgBase64 });
				$('#ipicker-modal').modal('hide');
			} else {
				window.alert('Please choose an image file.');
			}
		}
	}

	imgtoBase64 = file => new Promise((resolve, reject) => {
		const reader = new FileReader();
		reader.readAsDataURL(file);
		reader.onload = () => resolve(reader.result);
		reader.onerror = error => reject(error);
	});

	saveToServer = (data) => {
		let formData = new FormData();

		formData.append('imageData', data);
		formData.append('imageType', this.imageType);
		formData.append('uploadDir', this.config.uploadDir);
		formData.append('prefixFileName', this.config.prefixFileName);
		formData.append('imgSizes', JSON.stringify(this.config.imgSizes));

		/**
		 * Post data form to server using http
		 */
		let spinner = document.querySelector('.global-spinner');
		if (spinner) spinner.style.display = 'flex';
		let xhr = new XMLHttpRequest();
		xhr.open('POST', this.config.uploadUrl, true);
		xhr.onload = function (response) {
			if (xhr.readyState == XMLHttpRequest.DONE) {
				spinner.style.display = 'none';
				if (xhr.status === 200) {
					try {
						let res = JSON.parse(xhr.responseText);
						if (res.errorCode == 0 && res.data) {
							$(this.config.selector + ' .preview').append(this.createPreviewItem(res.data.fileName, res.data.url));
							this.updateInput(res.data.fileName);
							window.pixie.close();
						}
					} catch (e) { }
				} else {
					console.log("Error:", xhr.statusText);
				}
			}
		}.bind(this);
		xhr.send(formData);
	}

	createExistImgsPreview = () => {
		var value = $(this.config.selector + ' input').val();
		var urls = $(this.config.selector + ' input').attr('urls');
		if (value != '' && urls != '') {
			var fileNameArr = value.split(',');
			var urlsArr = urls.split(',');

			for (let i = 0; i < fileNameArr.length; i++) {
				$(this.config.selector + ' .preview').append(this.createPreviewItem(fileNameArr[i], urlsArr[i]));
			}
		}
	}

	createPreviewItem = (fileName, url) => {
		var previewItem = document.createElement('div');
		previewItem.className = 'preview-item';
		previewItem.innerHTML = `
			<div class="delete-image" title="Remove image">
				<i class="fa fa-trash"></i>
			</div>
			<img class="img img-responsive" src="${url}"/>
		`
		previewItem.querySelector('.delete-image').onclick = function () {
			let value = $(this.config.selector + ' input').val();
			if (value) {
				value = value.split(',');
			} else {
				value = [];
			}
			let index = value.indexOf(fileName);
			if (index > -1) {
				value.splice(index, 1);
			}
			$(this.config.selector + ' input').val(value.join(','));
			previewItem.remove();
		}.bind(this);

		return previewItem;
	}

	updateInput = (fileName) => {
		var value = $(this.config.selector + ' input').val();
		if (value) {
			value = value.split(',');
		} else {
			value = [];
		}
		value.push(fileName);

		$(this.config.selector + ' input').val(value.join(','));
	}
}