all: css/booklit.css css/slides.css

.PHONY: clean

clean:
	rm -f css/booklit.css
	rm -f css/slides.css

css/booklit.css: less/booklit.less
	yarn run lessc $< $@

css/slides.css: less/slides.less
	yarn run lessc $< $@
