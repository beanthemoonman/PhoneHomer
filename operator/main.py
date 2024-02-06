import json

import requests


def runAndPrint(response):
    print(response.text, response.status_code)


if __name__ == '__main__':
    def main():
        runAndPrint(requests.get(
            url="http://localhost:8080/getClientList",
            headers={"x-api-key": "1234"}
        ))
        runAndPrint(requests.post(
            url="http://localhost:8080/update",
            headers={"x-api-key": "1234"},
            data=json.dumps({
                'id': '12345678',
                'error': ''
            })
        ))
        runAndPrint(requests.get(
            url="http://localhost:8080/getClientList",
            headers={"x-api-key": "1234"}
        ))
        # for i in range(0, 19999):
        #     requests.post(
        #         url="http://localhost:8080/update",
        #         headers={"x-api-key": "1234"},
        #         data=json.dumps({
        #             'id': str(i),
        #             'error': ''
        #         })
        #     )
        #     print("post ", i)
        # runAndPrint(requests.get(
        #     url="http://localhost:8080/getClientList",
        #     headers={"x-api-key": "1234"}
        # ))
    main()
