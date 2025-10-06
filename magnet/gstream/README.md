Playing media straight from the Internet without storing it locally is known as Streaming. We have been doing it throughout the tutorials whenever we used a URI starting with http://. This tutorial shows a couple of additional points to keep in mind when streaming. In particular:


    * How to enable buffering (to alleviate network problems)
    * How to recover from interruptions (lost clock)

When streaming, media chunks are decoded and queued for presentation as soon as they arrive from the network. This means that if a chunk is delayed (which is not an uncommon situation at all on the Internet) the presentation queue might run dry and media playback could stall.



The universal solution is to build a “buffer”, this is, allow a certain number of media chunks to be queued before starting playback. In this way, playback start is delayed a bit, but, if some chunks are late, reproduction is not impacted as there are more chunks in the queue, waiting.
