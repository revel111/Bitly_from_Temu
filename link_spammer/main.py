import random
import string
import sys

import requests

def spammer(endpoint: str, times: int) -> None:
    print("Spam started")

    for i in range(times):
        print("Request number: ", i + 1)

        url = generate_url()
        payload = {"url": url}
        response = requests.post(endpoint, json=payload)
        print(f"Sent URL: {url}\nResponse Status: {response.status_code}\nResponse Body: {response.text}\n")

    print("Spam finished")

def generate_url () -> str:
    domain = "".join(random.choices(string.ascii_lowercase, k=8))
    path = "".join(random.choices(string.ascii_letters + string.digits, k=6))
    return f"https://{domain}.com/{path}"

# Provide entry point for script and amount of request  in arguments
if __name__ == "__main__":
    args = sys.argv

    if len(args) != 3:
        print("Wrong number of arguments!")
    else:
        spammer(args[1], int(args[2]))