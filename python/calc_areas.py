class Rectangle:
    def __init__(self, width, height):
        self.width = width
        self.height = height
        self.name = "Rectangle"
        
    def area(self):
        return self.width * self.height


class Circle:
    def __init__(self, radius):
        self.radius = radius
        self.name = "Circle"
        
        
    def area(self):
        return 3.14159 * self.radius ** 2


class Triangle:
    def __init__(self, base, height):
        self.base = base
        self.height = height
        self.name = "Triangle"
        
    def area(self):
        return 0.5 * self.base * self.height


def calc_area(shapes):
    for shape in shapes:
        print(f"The area of the {shape.name} is {shape.area()}")
        
        if hasattr(shape, "circumference"):
            print(f"The circumference of the {shape.name} is {shape.circumference()}")


shapes = [
    Rectangle(2, 3),
    Circle(5),
    Triangle(4, 6)
]

print("Areas of the shapes:")
calc_area(shapes)
