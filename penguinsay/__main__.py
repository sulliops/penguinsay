from linuxfacts import *
from text import *
from penguin import *
import random

def main():
    # Get random fact
    fact = random.choice(facts_list)
    # Output fact and penguin
    output(fact)
    penguin()

    return

if __name__ == "__main__":
    main()