# TEST CURL
import requests

if __name__ == '__main__':
    data = {"provider": "aliyun", "phone_number": ["", ""],
            "sign_name": "BlockFin", "template_code": "SMS_175485223", "template_param": {"code": "1234", "var": "123"}}
    header = "Content-type: application/json"
    re = requests.post("http://localhost:9876/alert", json=data)