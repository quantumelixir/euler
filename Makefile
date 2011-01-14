CC=6g
LD=6l

SOURCES=1.go 2.go
OBJECTS=$(SOURCES:.go=.6)
INCLUDE=-I/opt/local/lib/gopkg 
LIBRARY=-L/opt/local/lib/gopkg 

.SUFFIXES: .6 .go

.go.6:
	$(CC) $(INCLUDE) $<

.6:
	$(LD) $(LIBRARY) -o $* $<

all: $(SOURCES:.go=)

clean:
	rm -rf *.6
