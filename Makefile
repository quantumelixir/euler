all: install

DIRS=\
	 utils\
	 euler\

clean.dirs: $(addsuffix .clean, $(DIRS))
install.dirs: $(addsuffix .install, $(DIRS))

%.clean:
	+cd $* && gomake clean

%.install:
	+cd $* && gomake install

clean: clean.dirs

install: install.dirs

echo-dirs:
	@echo $(DIRS)
