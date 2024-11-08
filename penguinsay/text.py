# Creates textbox based on linux facts
def output(fact):
    text_length = len(fact) # Length of fact
    # Set up format for text box
    dashes = "  "
    text_str = "< " + fact + " >"

    for i in range(text_length):
        dashes += "-"

    # Output text box
    print(dashes)
    print(text_str)
    print(dashes)
    print(r"""        \
         \ """)
