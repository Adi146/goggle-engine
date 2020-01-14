import sys
import logging

from optparse import OptionParser
from pyassimp import load, postprocess
from struct import pack

logging.basicConfig(level=logging.DEBUG)

meshes = []


class Material:
    diffuse = []
    specular = []
    emissive = []
    shininess = 0.0

    def __init__(self, diffuse, specular, emissive, shininess):
        self.diffuse = diffuse
        self.specular = specular
        self.emissive = emissive
        self.shininess = shininess


class Mesh:
    name = ""
    vertices = []
    normals = []
    faces = []

    def __init__(self, name, vertices, normals, faces):
        self.name = name
        self.vertices = vertices
        self.normals = normals
        self.faces = faces


def process_mesh(mesh, scene):
    logging.debug("\t[" + str(mesh) + "] Start Processing Mesh...")

    vertices = []
    normals = []
    faces = []

    for vertex in mesh.vertices:
        vertices.append(vertex)
        
    for normal in mesh.normals:
        normals.append(normal)

    for face in mesh.faces:
        faces.append(face)

    material = process_material(scene.materials[mesh.materialindex], scene)

    meshes.append((Mesh(str(mesh), vertices, normals, faces), material))


def process_material(material, scene):
    logging.debug("\t\t[" + str(material) + "] Start Processing Material...")

    return Material(
        material.properties["diffuse"],
        material.properties["specular"],
        material.properties["emissive"],
        material.properties["shininess"]
    )


def process_node(node, scene):
    logging.debug("[" + str(node) + "] Start Processing Node...")

    for mesh in node.meshes:
        process_mesh(mesh, scene)

    for child in node.children:
        process_node(child, scene)


parser = OptionParser()
parser.add_option("-i", "--input", dest="input", help="path to model file")
parser.add_option("-o", "--output", dest="output", help="path to output file")
(options, args) = parser.parse_args()

scene = load(
    options.input, processing=
    postprocess.aiProcess_PreTransformVertices |
    postprocess.aiProcess_Triangulate |
    postprocess.aiProcess_GenNormals |
    postprocess.aiProcess_OptimizeMeshes |
    postprocess.aiProcess_OptimizeGraph |
    postprocess.aiProcess_JoinIdenticalVertices |
    postprocess.aiProcess_ImproveCacheLocality
)

if scene is None or scene.mRootNode is None:
    logging.error("Error while loading model")
    sys.exit(1)

process_node(scene.rootnode, scene)

logging.debug("Found " + str(len(meshes)) + " Meshes")
for mesh, material in meshes:
    logging.debug("Mesh: " + mesh.name)
    logging.debug("\tFound " + str(len(mesh.vertices)) + " Vertices")
    logging.debug("\tFound " + str(len(mesh.normals)) + " Normals")
    logging.debug("\tFound " + str(len(mesh.faces)) + " Faces")
    logging.debug("\tFound " + str(sum(len(x) for x in mesh.faces)) + " Indices")

    if len(mesh.vertices) != len(mesh.normals):
        logging.warning("Normals count does not match Vertex count")

logging.debug("Writing Output File...")
with open(options.output, "wb") as outFile:
    outFile.write(pack("<Q", len(meshes)))
    for mesh, material in meshes:
        outFile.write(pack("<QQ", len(mesh.vertices), sum(len(x) for x in mesh.faces)))

        for i in range(len(mesh.vertices)):
            outFile.write(pack("<" + str(len(mesh.vertices[i])) + "f", *mesh.vertices[i]))
            outFile.write(pack("<" + str(len(mesh.normals[i])) + "f", *mesh.normals[i]))

        for face in mesh.faces:
            outFile.write(pack("<" + str(len(face)) + "i", *face))

        outFile.write(pack("<" + str(len(material.diffuse)) + "f", *material.diffuse))
        outFile.write(pack("<" + str(len(material.specular)) + "f", *material.specular))
        outFile.write(pack("<" + str(len(material.emissive)) + "f", *material.emissive))
        outFile.write(pack("<f", material.shininess))

    outFile.close()
