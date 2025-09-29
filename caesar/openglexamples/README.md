* in Ubuntu use __apt-cache search libgl__ to search for the GL development files

  1. *GL* 
    
        OpenGL API implementation (http://www.opengl.org)
  2. *GLU* OpenGL Utility
  3. *Glut* 
  
        GLUT (OpenGL Utility Toolkit) – Glut is portable windowing API and it is not officially part of OpenGL.

        OpenGL Utility Toolkit (http://www.opengl.org/resources/libraries/glut/)


```sh
sudo apt-get install libglu1-mesa-dev freeglut3-dev mesa-common-dev
```

```sh
gcc cube.c -o cube $(pkg-config --libs glu glut gl)
```

* Examples

https://cs.lmu.edu/~ray/notes/openglexamples/

https://www.opengl.org/archives/resources/code/samples/glut_examples/examples/examples.html