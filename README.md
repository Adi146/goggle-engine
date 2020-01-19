# goggle-engine

GogGLE-Engine is a small OpenGL 3D scene graph written in GO. 
It can be used to create cross-platform GO applications which features interactive 3D representations.

Note: This is still in development. There might be bigger changes for the APIs!

# Compile on Windows:
1. Install mingw-w64
2. Download [SDL2](https://www.libsdl.org/download-2.0.php) development libraries for MinGW
3. For the modelconverter, you need to install following dependencies:
    * [assimp](https://packages.msys2.org/package/mingw-w64-x86_64-assimp)
    * [minizip](https://packages.msys2.org/package/mingw-w64-x86_64-minizip)
    * [zlib](https://packages.msys2.org/package/mingw-w64-x86_64-zlib)
4. Create a folder called "x86_64-w64-mingw32" in the mingw directory and copy the content of the archives in it.
5. Place the dlls in your execution path

## Packages
GogGLE consists of multiple Packages
### Core

Core handles all interactions with OpenGL and SDL.
You can use the core package by its own if you want to display a simple scene and don´t require the scene graph.

### SceneGraph 

Scene graph handles complex scenes where you can transform nodes depending on other nodes.
It is designed to be easily extensible with custom nodes.
The scene graph depends on core.

### UI

This package handles the User Input for interactive scenes.
It depends on the core and on the scene graph

### Examples

Here you can find some examples.

## Dependencies
* [go-gl](https://github.com/go-gl/gl)
* [sdl2](https://github.com/veandco/go-sdl2)
* [go-yaml v3](https://github.com/go-yaml/yaml)
* [assimp](https://github.com/Adi146/assimp)

