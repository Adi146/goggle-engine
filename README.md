# goggle-engine

GogGLE-Engine is a small OpenGL 3D scene graph written in GO. 
It can be used to create cross-platform GO applications which features interactive 3D representations.

Note: This is still in development. There might be bigger changes for the APIs!

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

