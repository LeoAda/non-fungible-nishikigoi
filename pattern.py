from randomizer import Randomizer
from random import choices
from constant import COLOR, COLOR_WEIGHT
from drawer import Drawer

if __name__ == '__main__':
    rando = Randomizer()
    choix = choices(COLOR, weights=COLOR_WEIGHT, k=6)

    border = 100
    nb_point = 10
    complexity = 50
    df = rando.rand_cords(nb_point,border)
    df = rando.cords_sorter(df)

    drawer = Drawer()
    drawer.shape_factory(border,df,complexity)


