import requests
import json


BASE_URL = "http://172.27.10.87:3001/api/v1" # 测试时需要修改
def create_comment(article_id, content, parent_comment_id=None):
    url = f"{BASE_URL}/comments"
    payload = {
        "article_id": article_id,
        "content": content,
        "parent_comment_id": parent_comment_id
    }
    response = requests.post(url, json=payload)
    
    if response.status_code == 200:
        print("评论创建成功！")
        print(json.dumps(response.json(), indent=2))
        return response.json()["comment"]["id"]
    else:
        print(f"创建评论失败: {response.status_code}")
        print(response.json())
        return None

def get_comments(article_id, sort_by='created_at', order='desc'):
    url = f"{BASE_URL}/articles/{article_id}/comments"
    params = {
        "sort_by": sort_by,
        "order": order
    }
    response = requests.get(url, params=params)
    
    if response.status_code == 200:
        print("评论列表：")
        print(json.dumps(response.json(), indent=2))
    else:
        print(f"获取评论失败: {response.status_code}")
        print(response.json())

def delete_comment(comment_id):
    url = f"{BASE_URL}/comments/{comment_id}"
    response = requests.delete(url)
    
    if response.status_code == 200:
        print(f"评论 {comment_id} 删除成功！")
    else:
        print(f"删除评论失败: {response.status_code}")
        print(response.json())

def test_create_delete_comments():
    article_id = 1

    root_comment_id = create_comment(article_id, "这是根评论")
    if not root_comment_id:
        return

    child_comment_id = create_comment(article_id, "这是回复根评论的子评论", parent_comment_id=root_comment_id)
    if not child_comment_id:
        return

    print("获取评论列表（创建根评论和子评论后）：")
    get_comments(article_id)

    sub_child_comment_id = create_comment(article_id, "这是回复子评论的孙评论", parent_comment_id=child_comment_id)
    if not sub_child_comment_id:
        return

    print("获取评论列表（创建三级评论后）：")
    get_comments(article_id)

    delete_comment(root_comment_id)

    print("获取评论列表（删除根评论后）：")
    get_comments(article_id)

if __name__ == "__main__":
    test_create_delete_comments()