Tea.context(function () {
	var that = this;

	this.items.$each(function (k, v) {
		v.isChecked = that.chart.requirements.$contains(v.code);
	});

	this.$delay(function () {
		this.$find("form input[name='name']").focus();
		this.loadJavascriptChart();
	});

	this.submitSuccess = function () {
		alert("保存成功");
		window.location = this.from;
	};

	/**
	 * 指标
	 */
	this.uncheckItem = function (item) {
		item.isChecked = false;
	};

	/**
	 * 更多选项
	 */
	this.advancedOptionsVisible = false;
	this.showAdvancedOptions = function () {
		this.advancedOptionsVisible = !this.advancedOptionsVisible;
	};

	/**
	 * Javascript chart
	 */
	this.loadJavascriptChart = function () {
		var editor = CodeMirror.fromTextArea(document.getElementById("javascript-code-editor"), {
			theme: "idea",
			lineNumbers: true,
			value: "",
			readOnly: false,
			showCursorWhenSelecting: true,
			height: "auto",
			//scrollbarStyle: null,
			viewportMargin: Infinity,
			lineWrapping: true,
			highlightFormatting: false,
			indentUnit: 4,
			indentWithTabs: true
		});

		editor.setValue(this.chart.options.code);
		editor.save();

		var info = CodeMirror.findModeByMIME("text/javascript");
		if (info != null) {
			editor.setOption("mode", info.mode);
			CodeMirror.modeURL = "/codemirror/mode/%N/%N.js";
			CodeMirror.autoLoadMode(editor, info.mode);
		}

		var that = this;
		editor.on("change", function () {
			editor.save();
			that.chartCode = editor.getValue();
		});

		//editor.focus();
	};

	/**
	 * 测试
	 */
	this.charts = [];
	this.isLoaded = false;
	this.intervalSeconds = 5;

	this.test = function () {
		var form = this.$find("#make-form")[0];
		var formData = new FormData(form);
		this.$post("/proxy/board/test")
			.params(formData)
			.success(function (resp) {
				// output
				resp.data.output.$each(function (k, v) {
					console.log("[widget]" + v);
				});

				// charts
				this.charts = resp.data.charts;
				new ChartRender(this.charts);

				this.$delay(function () {
					window.scroll(0, 10000);
				}, 100);
			})
			.done(function () {
				this.isLoaded = true;
			});
	};
});