package engine

func init() {
	DeclFunc("Flush", drainOutput, "Flush all pending output to disk.")
	DeclFunc("OldAutoSave", AutoSave, "Auto save space-dependent quantity every period (s).")
	DeclFunc("AutoSnapshot", AutoSnapshot, "Auto save image of quantity every period (s).")
	DeclFunc("Chunk", Mx3chunks, "")
	DeclFunc("Uniform", Uniform, "Uniform magnetization in given direction")
	DeclFunc("Vortex", Vortex, "Vortex magnetization with given circulation and core polarization")
	DeclFunc("Antivortex", AntiVortex, "Antivortex magnetization with given circulation and core polarization")
	DeclFunc("Radial", Radial, "Radial magnetization with given charge and core polarization")
	DeclFunc("NeelSkyrmion", NeelSkyrmion, "Néél skyrmion magnetization with given charge and core polarization")
	DeclFunc("BlochSkyrmion", BlochSkyrmion, "Bloch skyrmion magnetization with given chirality and core polarization")
	DeclFunc("TwoDomain", TwoDomain, "Twodomain magnetization with with given magnetization in left domain, wall, and right domain")
	DeclFunc("VortexWall", VortexWall, "Vortex wall magnetization with given mx in left and right domain and core circulation and polarization")
	DeclFunc("RandomMag", RandomMag, "Random magnetization")
	DeclFunc("RandomMagSeed", RandomMagSeed, "Random magnetization with given seed")
	DeclFunc("Conical", Conical, "Conical state for given wave vector, cone direction, and cone angle")
	DeclFunc("Helical", Helical, "Helical state for given wave vector")
	DeclFunc("Crop", Crop, "Crops a quantity to cell ranges [x1,x2[, [y1,y2[, [z1,z2[")
	DeclFunc("CropX", CropX, "Crops a quantity to cell ranges [x1,x2[")
	DeclFunc("CropY", CropY, "Crops a quantity to cell ranges [y1,y2[")
	DeclFunc("CropZ", CropZ, "Crops a quantity to cell ranges [z1,z2[")
	DeclFunc("CropLayer", CropLayer, "Crops a quantity to a single layer")
	DeclFunc("CropRegion", CropRegion, "Crops a quantity to a region")

	DeclFunc("AddFieldTerm", AddFieldTerm, "Add an expression to B_eff.")
	DeclFunc("AddEdensTerm", AddEdensTerm, "Add an expression to Edens.")
	DeclFunc("Add", Add, "Add two quantities")
	DeclFunc("Madd", Madd, "Weighted addition: Madd(Q1,Q2,c1,c2) = c1*Q1 + c2*Q2")
	DeclFunc("Dot", Dot, "Dot product of two vector quantities")
	DeclFunc("Cross", Cross, "Cross product of two vector quantities")
	DeclFunc("Mul", Mul, "Point-wise product of two quantities")
	DeclFunc("MulMV", MulMV, "Matrix-Vector product: MulMV(AX, AY, AZ, m) = (AX·m, AY·m, AZ·m). "+
		"The arguments Ax, Ay, Az and m are quantities with 3 componets.")
	DeclFunc("Div", Div, "Point-wise division of two quantities")
	DeclFunc("Const", Const, "Constant, uniform number")
	DeclFunc("ConstVector", ConstVector, "Constant, uniform vector")
	DeclFunc("Shifted", Shifted, "Shifted quantity")
	DeclFunc("Masked", Masked, "Mask quantity with shape")
	DeclFunc("Normalized", Normalized, "Normalize quantity")
	DeclFunc("RemoveCustomFields", RemoveCustomFields, "Removes all custom fields again")

	DeclFunc("ext_ScaleExchange", ScaleInterExchange, "Re-scales exchange coupling between two regions.")
	DeclFunc("ext_InterExchange", InterExchange, "Sets exchange coupling between two regions.")
	DeclFunc("ext_ScaleDind", ScaleInterDind, "Re-scales Dind coupling between two regions.")
	DeclFunc("ext_InterDind", InterDind, "Sets Dind coupling between two regions.")
	DeclFunc("ext_centerBubble", CenterBubble, "centerBubble shifts m after each step to keep the bubble position close to the center of the window")
	DeclFunc("ext_centerWall", CenterWall, "centerWall(c) shifts m after each step to keep m_c close to zero")
	DeclFunc("ext_make3dgrains", Voronoi3d, "3D Voronoi tesselation over shape (grain size, starting region number, num regions, shape, seed)")
	DeclFunc("ext_makegrains", Voronoi, "Voronoi tesselation (grain size, num regions)")
	DeclFunc("ext_rmSurfaceCharge", RemoveLRSurfaceCharge, "Compensate magnetic charges on the left and right sides of an in-plane magnetized wire. Arguments: region, mx on left and right side, resp.")
	DeclFunc("SetGeom", SetGeom, "Sets the geometry to a given shape")
	DeclFunc("ReCreateMesh", ReCreateMesh, "")
	DeclFunc("Minimize", Minimize, "Use steepest conjugate gradient method to minimize the total energy")

	DeclFunc("DefRegion", DefRegion, "Define a material region with given index (0-255) and shape")
	DeclFunc("ShapeFromRegion", ShapeFromRegion, "")
	DeclFunc("DefRegionCell", DefRegionCell, "Set a material region (first argument) in one cell "+
		"by the index of the cell (last three arguments)")
	DeclFunc("Relax", Relax, "Try to minimize the total energy")
	DeclFunc("Run", Run, "Run the simulation for a time in seconds")
	DeclFunc("RunWithoutPrecession", RunWithoutPrecession, "Run the simulation for a time in seconds with precession disabled")
	DeclFunc("Steps", Steps, "Run the simulation for a number of time steps")
	DeclFunc("RunWhile", RunWhile, "Run while condition function is true")
	DeclFunc("SetSolver", SetSolver, "Set solver type. 1:Euler, 2:Heun, 3:Bogaki-Shampine, 4: Runge-Kutta (RK45), 5: Dormand-Prince, 6: Fehlberg, -1: Backward Euler")
	DeclFunc("Exit", Exit, "Exit from the program")
	DeclFunc("RunShell", runShell, "Run a shell command")

	DeclFunc("SaveOvf", Save, "Save space-dependent quantity once, with auto filename")
	DeclFunc("SaveOvfAs", SaveAs, "Save space-dependent quantity with custom filename")
	DeclFunc("Snapshot", Snapshot, "Save image of quantity")
	DeclFunc("SnapshotAs", SnapshotAs, "Save image of quantity with custom filename")

	DeclFunc("Ellipsoid", Ellipsoid, "3D Ellipsoid with axes in meter")
	DeclFunc("Ellipse", Ellipse, "2D Ellipse with axes in meter")
	DeclFunc("Cone", Cone, "3D Cone with diameter and height in meter. The base is at z=0. If the height is positive, the tip points in the +z direction.")
	DeclFunc("Cylinder", Cylinder, "3D Cylinder with diameter and height in meter")
	DeclFunc("Circle", Circle, "2D Circle with diameter in meter")
	DeclFunc("Squircle", Squircle, "2D Squircle with diameter in meter")
	DeclFunc("Cuboid", Cuboid, "Cuboid with sides in meter")
	DeclFunc("Rect", Rect, "2D rectangle with size in meter")
	DeclFunc("Wave", Wave, "Wave with (Period, Min amplitude and Max amplitude) in meter")
	DeclFunc("Triangle", Triangle, "Equilateral triangle with side in meter")
	DeclFunc("RTriangle", RTriangle, "Rounded Equilateral triangle with side in meter")
	DeclFunc("Diamond", Diamond, "Diamond with side in meter")
	DeclFunc("Hexagon", Hexagon, "Hexagon with side in meter")
	DeclFunc("Square", Square, "2D square with size in meter")
	DeclFunc("XRange", XRange, "Part of space between x1 (inclusive) and x2 (exclusive), in meter")
	DeclFunc("YRange", YRange, "Part of space between y1 (inclusive) and y2 (exclusive), in meter")
	DeclFunc("ZRange", ZRange, "Part of space between z1 (inclusive) and z2 (exclusive), in meter")
	DeclFunc("Layers", Layers, "Part of space between cell layer1 (inclusive) and layer2 (exclusive), in integer indices")
	DeclFunc("Layer", Layer, "Single layer (along z), by integer index starting from 0")
	DeclFunc("Universe", Universe, "Entire space")
	DeclFunc("Cell", Cell, "Single cell with given integer index (i, j, k)")
	DeclFunc("ImageShape", ImageShape, "Use black/white image as shape")
	DeclFunc("GrainRoughness", GrainRoughness, "Grainy surface with different heights per grain "+
		"with a typical grain size (first argument), minimal height (second argument), and maximal "+
		"height (third argument). The last argument is a seed for the random number generator.")
	DeclFunc("Shift", Shift, "Shifts the simulation by +1/-1 cells along X")
	DeclFunc("ThermSeed", ThermSeed, "Set a random seed for thermal noise")

	DeclFunc("Expect", Expect, "Used for automated tests: checks if a value is close enough to the expected value")
	DeclFunc("ExpectV", ExpectV, "Used for automated tests: checks if a vector is close enough to the expected value")
	DeclFunc("Fprintln", Fprintln, "Print to file")
	DeclFunc("Sign", sign, "Signum function")
	DeclFunc("Vector", Vector, "Constructs a vector with given components")
	DeclFunc("Print", myprint, "Print to standard output")
	DeclFunc("LoadFile", LoadFile, "Load a zarr data file")
	DeclFunc("LoadOvfFile", LoadOvfFile, "Load an ovf data file")
	DeclFunc("Index2Coord", Index2Coord, "Convert cell index to x,y,z coordinate in meter")
	DeclFunc("NewSlice", NewSlice, "Makes a 4D array with a specified number of components (first argument) "+
		"and a specified size nx,ny,nz (remaining arguments)")
	DeclFunc("NewVectorMask", NewVectorMask, "Makes a 3D array of vectors")
	DeclFunc("NewScalarMask", NewScalarMask, "Makes a 3D array of scalars")
	DeclFunc("RegionFromCoordinate", RegionFromCoordinate, "RegionFromCoordinate")

	DeclFunc("AutoSaveAs", Mx3AutoSaveAs, "Auto save space-dependent quantity every period (s) as the zarr standard.")
	DeclFunc("AutoSaveAsChunk", Mx3AutoSaveAsChunk, "Auto save space-dependent quantity every period (s) as the zarr standard.")
	DeclFunc("AutoSave", Mx3AutoSave, "Auto save space-dependent quantity every period (s) as the zarr standard.")
	DeclFunc("SaveAs", Mx3SaveAs, "Save space-dependent quantity as the zarr standard.")
	DeclFunc("SaveAsChunk", Mx3SaveAsChunk, "")
	DeclFunc("Save", Mx3zSave, "Save space-dependent quantity as the zarr standard.")

	DeclFunc("TableSave", TableSave, "Save the data table right now.")
	DeclFunc("TableAdd", TableAdd, "Save the data table periodically.")
	DeclFunc("TableAddVar", TableAddVar, "Save the data table periodically.")
	DeclFunc("TableAddAs", TableAddAs, "Save the data table periodically.")
	DeclFunc("TableAutoSave", TableAutoSave, "Save the data table periodically.")
}

// Add a function to the script world
func DeclFunc(name string, f interface{}, doc string) {
	World.Func(name, f, doc)
}
