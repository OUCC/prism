from settings import *

from datetime import datetime
from urllib.request import urlopen
from urllib.parse import urlencode
from urllib.error import HTTPError, URLError
import base64
import time
import os


def post(image):
    b64 = base64.b64encode(image)
    data = {'image': b64, 'key': POST_KEY}

    try:
        urlopen(url=POST_URL, data=urlencode(data).encode())

    except HTTPError as e:
        print("The server couldn't fulfill the request.")
        print('Error code: ', e.code)
        return False

    except URLError as e:
        print('We failed to reach a server.')
        print('Reason: ', e.reason)
        return False

    return True

    
if __name__ == "__main__":
    while True:
        # capture image and save
        os.system('python2 '+os.path.dirname(os.path.abspath(__file__))+'/capture/capture.py')

        with open('tmp.jpg', 'rb') as img:
            if post(img.read()):
                print(datetime.today().strftime(TIME_FORMAT) + " Post success")
            else:
                print(datetime.today().strftime(TIME_FORMAT) + " Post failed")

        # capturing takes about 3sec...
        time.sleep(297)
