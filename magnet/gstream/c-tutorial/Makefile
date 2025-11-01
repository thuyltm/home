CXX=g++
CXXFLAGS=-g -Wall `pkg-config --cflags --libs gstreamer-1.0 gstreamer-video-1.0`
SOURCES=$(wildcard *.c *cpp)
TARGET=$(SOURCES:.c= $(basename $$SOURCES))
OBJECTS=$(TARGET:=.o)
clean:
	rm -f $(TARGET) $(OBJECTS)
# Pattern rule to compile each .c file into a .o file
%o: %.c
	$(CC) $< -o $@ $(CXXFLAGS) 
%: %.c
	$(CC) $< -o $@ $(CXXFLAGS)