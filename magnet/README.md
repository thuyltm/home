* Install Gstream

https://gstreamer.freedesktop.org/documentation/installing/on-linux.html?gi-language=c

```sh
pkg-config --cflags --libs gstreamer-1.0
```

* Install FFMPEG
    * libavcodec: Contains decoders and encoders for a wide range of audio and video codecs.
    * libavformat: Implements streaming protocols, container formats, and basic I/O access for various multimedia formats. 
    * libavutil: Provides utility functions, including hashers, decompressors, and other miscellaneous tools.
    * libavfilter: Offers a powerful framework for altering decoded audio and video through a directed graph of connected filters.
    * libavdevice: Provides an abstraction layer for accessing capture and playback devices.
    * libswresample: Implements routines for audio mixing and resampling.
    * libswscale: Handles color conversion and scaling routines for video.

### Compare ffmpeg vs gstream
* FFmpeg:
    * Nature: Primarily __a command-line tool__ and a set of libraries for __handling audio and video__. It's renowned for its robust __encoding and decoding capabilities__, extensive codec support, and versatility in various multimedia tasks.
    * Strengths: Excellent for transcoding, format conversion, basic video manipulation (like cropping, scaling), and general multimedia processing via its command-line interface. It's a go-to for quick and efficient media operations
    * Limitations: While powerful, its command-line nature can be __less intuitive for complex, real-time streaming__ or pipeline-based applications.


* GStreamer:

    * Nature: __A pipeline-based multimedia framework__ designed for building __complex media processing graphs__. It uses a modular, plugin-based architecture, allowing developers to create custom pipelines __by connecting various "elements" (plugins)__ that perform specific tasks.
    * Strengths: Ideal for __building sophisticated multimedia applications, real-time streaming,__ intricate audio/video processing chains, and scenarios requiring fine-grained control over the media flow. Its modularity and pipeline design make it highly flexible and extensible.
    * Limitations: Can have a steeper learning curve due to its framework nature and the __need to understand pipeline construction__. For simple, one-off media conversions, it might be __overkill compared to FFmpeg__.