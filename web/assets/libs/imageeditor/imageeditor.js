'use strict';

const REGEXP_MIME_TYPE_IMAGES = /^image\/\w+$/;

const TOOLBAR_CROP_HTML = `
<div class="btn-group">
	<button type="button" class="btn btn-primary btn-crop" data-method="setDragMode" data-option="crop" title="Crop">
		<span class="docs-tooltip" data-toggle="tooltip" title="this.cropper.setDragMode(&quot;crop&quot;)">
			<span class="fa fa-crop-alt"></span>
		</span>
	</button>
</div>`;

const TOOLBAR_ROTATE_HTML = `
<div class="btn-group">
	<button type="button" class="btn btn-primary btn-rotate-left" data-method="rotate" data-option="-45" title="Rotate Left">
		<span class="docs-tooltip" data-toggle="tooltip" title="this.cropper.rotate(-45)">
			<span class="fa fa-undo-alt"></span>
		</span>
	</button>
	<button type="button" class="btn btn-primary btn-rotate-right" data-method="rotate" data-option="45" title="Rotate Right">
		<span class="docs-tooltip" data-toggle="tooltip" title="this.cropper.rotate(45)">
			<span class="fa fa-redo-alt"></span>
		</span>
	</button>
</div>`;

function createCropRatiosHTML(cropRatios = []) {
	const MAX = 40;
	var html = '<div class="panel-group crop-ratios">';
	(cropRatios || []).forEach(ratio => {
		const sizeArr = ratio.split(':');
		if (sizeArr.length == 2) {
			const width = parseFloat(sizeArr[0]);
			const height = parseFloat(sizeArr[1]);
			var previewWidth = 0;
			var previewHeight = 0;
			if (width / height > 1) {
				previewWidth = MAX;
				previewHeight = MAX * height / width;
			} else {
				previewWidth = MAX * width / height;
				previewHeight = MAX;
			}
			html += `<label class="crop-ratio">
				<div class="ratio-container">
					<input type="radio" class="sr-only" id="aspectRatio1" name="aspectRatio" value="${width / height}" />
					<div class="ratio-preview" style="width:${previewWidth}px;height:${previewHeight}px;"></div>
					<span class="docs-tooltip" data-toggle="tooltip" title="aspectRatio: ${width} / ${height}">
						${ratio}
					</span>
				</div>
			</label>`;
		}
	});
	// Add Free ratio option
	html += `<label class="crop-ratio free-ratio">
		<div class="ratio-container">
			<input type="radio" class="sr-only" id="aspectRatio5" name="aspectRatio" value="NaN" />
			<div class="ratio-preview"></div>
			<span class="docs-tooltip" data-toggle="tooltip" title="aspectRatio: NaN">
				Free
			</span>
		</div>
	</label>`;
	html += `<div class="actions">
		<button type="button" class="btn btn-outline-secondary cancel" title="Cancel crop">
			<span class="docs-tooltip" data-toggle="tooltip" title="Cancel crop">
				<span class="fa fa-window-close"></span>
			</span>
		</button>
		<button type="button" class="btn btn-outline-success apply" title="Crop">
			<span class="docs-tooltip" data-toggle="tooltip" title="Crop">
				<span class="fa fa-check"></span>
			</span>
		</button>
	</div>`;
	html += '</div>';
	return html;
}

function readFile(file, event) {
	return new Promise((resolve, reject) => {
		if (!file) {
			resolve();
			return;
		}
		if (REGEXP_MIME_TYPE_IMAGES.test(file.type)) {
			if (URL) {
				resolve({
					loaded: true,
					name: file.name,
					type: file.type,
					url: URL.createObjectURL(file),
				});
			} else {
				reject(new Error('Your browser is not supported.'));
			}
		} else {
			reject(new Error(`Please ${event ? event.type : 'choose'} an image file.`));
		}
	});
}

/**
 * 
 * @param {*} selector
 * @param {*} options
 * - toolbar:
 * -- crop: boolean (show/hide crop button)
 * -- cropRatios: []
 * -- rotate: boolean (show/hide rotate buttons)
 * - cropperOptions: Cropper js options, get from cropperjs library
 */
function ImageEditor(selector, options) {
	const image = $(selector).find('.image-source')[0];
	if (!image) {
		throw 'Not found image source';
	}
	
	this.cropper = new Cropper(image, {
		autoCrop: false,
		dragMode: 'none',
		viewMode: 2,
		autoCropArea: 1,
		ready: () => {
			if (options && options.toolbar) {
				const toolbarOptions = options.toolbar;
				const $toolbarSelector = $(selector).find('.image-editor-toolbar');
				$toolbarSelector.empty();
				$toolbarSelector.append('<div class="toolbar-buttons"></div>');
				$toolbarSelector.append('<div class="toolbar-panels d-none"></div>');

				const $toolbarButtons = $(selector).find('.toolbar-buttons');
				const $toolbarPanels = $(selector).find('.toolbar-panels');

				/**
				 * Toolbar buttons
				 */
				if (toolbarOptions.crop) {
					$toolbarButtons.append(TOOLBAR_CROP_HTML);
					$toolbarButtons.find('.btn-crop').click(() => {
						this.showCropPanel();
						// Set default crop to first ratio
						$toolbarPanels.find('.crop-ratios .crop-ratio').first().click();
					});
				}
				if (toolbarOptions.rotate) {
					$toolbarButtons.append(TOOLBAR_ROTATE_HTML);
					$toolbarButtons.find('.btn-rotate-left').click(() => this.rotate(-45));
					$toolbarButtons.find('.btn-rotate-right').click(() => this.rotate(45));
				}

				/**
				 * Toolbar panel
				 */
				if (toolbarOptions.crop && (toolbarOptions.cropRatios || []).length > 0) {
					$toolbarPanels.append(createCropRatiosHTML(toolbarOptions.cropRatios));

					const $cropRatios = $toolbarPanels.find('.crop-ratios');
					$cropRatios.find('.crop-ratio').click((e) => {
						this.setAspectRatio(e.target);
					});

					$cropRatios.find('.actions .cancel').click(this.cancelCrop);
					$cropRatios.find('.actions .apply').click(() => {
						this.cropImage();
					});
				}
			}
		}
	});

	this.showToolbarButtons = () => {
		$(selector).find('.toolbar-buttons').removeClass('d-none');
	}

	this.hideToolbarButtons = () => {
		$(selector).find('.toolbar-buttons').addClass('d-none');
	}

	this.showToolbarPanels = () => {
		$(selector).find('.toolbar-panels').removeClass('d-none');
	}

	this.hideToolbarPanels = () => {
		$(selector).find('.toolbar-panels').addClass('d-none');
	}

	this.showCropPanel = () => {
		const $toolbarPanels = $(selector).find('.toolbar-panels');
		$toolbarPanels.find('.panel-group').addClass('d-none');
		$toolbarPanels.find('.panel-group.crop-ratios').removeClass('d-none');
		this.hideToolbarButtons();
		this.showToolbarPanels();
	}

	this.rotate = (d) => {
		this.cropper.rotate(d);
	}

	this.setAspectRatio = (el) => {
		$(selector).find('.toolbar-panels .crop-ratios .crop-ratio').removeClass('active');
		$(el).addClass('active');
		const ratio = $(el).find('input[name="aspectRatio"]').val();
		this.cropper.setAspectRatio(ratio);
		this.cropper.crop();
	}

	this.cancelCrop = () => {
		this.cropper.clear();
		this.hideToolbarPanels();
		this.showToolbarButtons();
	}

	this.cropImage = () => {
		const canvas = this.getImageCanvasData();
		this.cropper.replace(canvas.toDataURL('image/jpeg'));
		this.cancelCrop();
	}

	this.getImageCanvasData = () => {
		const canvas = this.cropper.getCroppedCanvas({
			imageSmoothingEnabled: false,
			imageSmoothingQuality: 'high',
		});
		return canvas;
	}

	this.getImageDataURL = (mimeType = 'image/jpeg') => {
		const canvas = this.getImageCanvasData();
		return canvas.toDataURL(mimeType);
	}

	this.destroy = () => {
		if (this.cropper) {
			this.cropper.destroy();
		}
	}
}