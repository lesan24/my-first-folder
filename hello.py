import datetime

print("I am learning Computer Technology")
print(f"Today is {datetime.datetime.now()}")
print("My phone CPU is aarch64 - 64bit power!")

# Let's read the file you made earlier
with open("hello.txt", "r") as f:
    print("\n--- Content of hello.txt ---")
    print(f.read())
