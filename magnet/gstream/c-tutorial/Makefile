CXX=g++
CXXFLAGS=-g -Wall `pkg-config --cflags gstreamer-1.0 gstreamer-video-1.0`
LDFLAGS=`pkg-config --libs gstreamer-1.0 gstreamer-video-1.0`	
TARGET=tutorial-1 tutorial-2 tutorial-3 tutorial-4 tutorial-5 tutorial-6 tutorial-7 tutorial-8 tutorial-9 tutorial-10 tutorial-11 tutorial-12 tutorial-13 tutorial-14
OBJECTS=$(TARGET:=.o)
SOURCES=$(TARGET:=.cpp)
.PHONY: all clean
all: $(TARGET)
$(TARGET): %: %.o
	$(CXX) -o $@ $< $(LDFLAGS)
.cpp.o:
	$(CXX) -c $< $(CXXFLAGS)
clean:
	rm -f $(TARGET) $(OBJECTS)	
	