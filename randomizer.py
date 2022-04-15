from random import seed, random, randint,choice
from constant import COLOR, COLOR_WEIGHT
import pandas as pd
import numpy as np


class Randomizer:
    """
    This class is used to gather function that are used for random purposes.

    Attributes
    ----------

    Methods
    -------
    rand_cords(nb_point,border)
        Generate random integer coordinate.
    cords_sorter(df)
        Sort a dataframe of coordinates by argument.
    """

    def __init__(self, seeds=None):
        """
        Parameters
        ----------
        seeds : int, optional
            The seed of all random functions (default is None)"""
        if seeds is not None: seed(seeds)

    def rand_cords(self, nb_point, border):
        """
        Generate random integer coordinate between -border and border.
        Parameters
        ----------
        nb_point : int
            Number of points that it generates.
        border : int
            Max and min values.
        Returns
        -------
        pandas.Dataframe
            A dataframe with columns x,y and random integers as row.
        """
        lx, ly = [], []
        [lx.append(randint(-border, border)) for x in range(nb_point)]
        [ly.append(randint(-border, border)) for y in range(nb_point)]
        return pd.DataFrame({"x": lx, "y": ly})

    def cords_sorter(self, df):
        """
        Sort a dataframe of coordinates by argument.
        Parameters
        ----------
        df : pandas.DataFrame
            Pandas DataFrame with a column x and a column y
        Returns
        -------
        pandas.Dataframe
            Input DataFrame sorted with a new column theta (Arg)
        """
        thetas = np.arctan2(df.y, df.x)
        df.insert(2, "theta", thetas)
        df.loc[df['theta'] < 0, 'theta'] += 2 * np.pi  # Normalisation for negative arguments
        df = df.sort_values("theta").reset_index(drop=True)
        return df

    def rand_neighborhood(self,x,y,sigma):
        x1 = randint(x-sigma,x+sigma)
        y1 = randint(y-sigma,y+sigma)
        return x1,y1
