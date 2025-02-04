class Difference:
    def __init__(self, a):
        self.__elements = a

    maximumDifference = 0

    # Add your code here
    def computeDifference(self):
        self.__elements.sort()
        min = self.__elements[0]
        max = self.__elements[-1]
        self.maximumDifference = max - min

# End of Difference class

_ = input()
a = [int(e) for e in input().split(' ')]

d = Difference(a)
d.computeDifference()

print(d.maximumDifference)
