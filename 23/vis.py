from PIL import Image

f = open("./out.txt").read().strip().split("\n\n")[:-1]

colorDict = {
	'.': (0, 0, 0, 0),
	'#': (255, 255, 255, 255),
	'^': (128, 255, 0, 255),
	'v': (128, 0, 255, 255),
	'<': (0, 255, 255, 255),
	'>': (255, 0, 0, 255),
}

(size, f) = ([int(x) for x in f[0].split(" ")], f[1:])
f = [v.split("\n")[1:] if v.startswith("step") else v.split("\n") for v in f]
f = [(tuple(map(int, v[0].strip("[").split(" ")[:3])), v[1:]) for v in f]
f = [((s[0], s[2]), v) for (s, v) in f]
for n in range(len(f)):
	i = Image.new("RGB", size)
	px = i.load()
	yo, xo = f[n][0]
	for yi, row in enumerate(f[n][1]):
		for xi, v in enumerate(row):
			px[xi+xo, yi+yo] = colorDict[v]
	i = i.resize((size[0]*6, size[1]*6), Image.NEAREST)
	i.save(f"out/{n:04d}.png")
