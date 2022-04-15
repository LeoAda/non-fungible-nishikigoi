import svgwrite
from randomizer import Randomizer

class Drawer:
    """
    This class is used to gather function that are used for drawing and managing svg.

    Attributes
    ----------
    file_name : str, optional
            Filename of the output file (default is "basic_svg")
    Methods
    -------
    set_file_name(str new_file_name)
        Change name of the following file

    basic_shapes()
        description_of_method
    """

    def __init__(self, file_name="basic_svg"):
        """
        Parameters
        ----------
        file_name : str, optional
            Filename of the output file (default is "basic_svg")
        randomizer : Randomizer
            Instance of a randomizer
        """
        self.file_name = file_name
        self.randomizer = Randomizer()

    def set_file_name(self, new_file_name):
        """
        Setter of file_name
        Parameters
        ----------
        new_file_name : str
            New name of the following file create in svg.
        """
        self.file_name = new_file_name

    def basic_shapes(self):
        """
        Create a basic shape in svg
        """
        dwg = svgwrite.Drawing(filename=self.file_name + ".svg", debug=True)
        shapes = dwg.add(dwg.g(id='shapes'))
        path = "m 0 0 C 0 -45 138 -35 80 0 C 49 28 175 42 114 70 C -14 81 55 24 0 0"
        shapes.add(
            dwg.path(d=path, transform="translate(50,50)"))
        dwg.save()

    def shape_factory(self, border, df,sigma):
        dwg = svgwrite.Drawing(filename=self.file_name + ".svg", debug=True)
        shapes = dwg.add(dwg.g(id='shapes'))
        path = self.path_shape_generator(df,sigma)
        print(path)
        shapes.add(
            dwg.path(d=path, transform=f"translate({border},{border})"))
        dwg.save()

    def path_shape_generator(self, df,sigma):
        """
        Create a shape of path from a dataframe of points.
        Parameters
        ----------
        df : pandas.DataFrame
            A sorted dataframe of coordinates 'x' and 'y'
        Returns
        -------
        str
            A formatted path string of svg.
        """
        paths = ""
        for index, row in df.iterrows():
            a = 'Z'
            if index == 0:
                paths += f"m {row['x']} {row['y']} "
            else:
                xn1, yn1 = self.randomizer.rand_neighborhood(df.iloc[index-1].x,df.iloc[index-1].y,sigma)
                xn2, yn2 = self.randomizer.rand_neighborhood(df.iloc[index - 1].x, df.iloc[index - 1].y,sigma)
                paths += f"C {xn1} {yn1} {xn2} {yn2} {row['x']} {row['y']} "
        xn1, yn1 = self.randomizer.rand_neighborhood(df.iloc[-1].x, df.iloc[-1].y,sigma)
        xn2, yn2 = self.randomizer.rand_neighborhood(df.iloc[0].x, df.iloc[0].y,sigma)
        paths += f"C {xn1} {yn1} {xn2} {yn2} {df.iloc[0].x} {df.iloc[0].y}"  # Back to beginning
        return paths
