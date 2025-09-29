CXXFLAGS=-g -Wall `pkg-config --cflags glu glut gl`
LDFLAGS=`pkg-config --libs glu glut gl`
SOURCES=$(wildcard *.c *cpp)
OBJECTS=$(SOURCES:.c= $(basename $$SOURCES))
.PHONY: all clean
cube: cube.c
	$(CXX) $< $(CXXFLAGS) $(LDFLAGS) -o $@
run_cube: cube
	./cube
trackball.o: trackball.c trackball.h
	$(CXX) -c $^ $(CXXFLAGS) -o $@
clean:
	rm -f $(OBJECTS)
compile:
	@echo $(OBJECTS)
	for source in $(OBJECTS); do \
		$(CXX)  $(addsuffix .c, $$source) $(CXXFLAGS) $(LDFLAGS) -o $$source;\
	done
run:
	for source in $(OBJECTS); do \
		./$$source; \
	done
all: compile run
test:
	@echo $(SOURCES)