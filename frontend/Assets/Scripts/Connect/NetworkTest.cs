
using System.Collections;
using UnityEngine.Networking;
using UnityEngine;

public class NetworkTest : MonoBehaviour
{
    // Start is called before the first frame update
    void Start()
    {
        StartCoroutine(Post());
    }

    protected IEnumerator Post()
    {
        WWWForm form = new WWWForm();
        // 使用 AddField 帶入參數
        string URL = "http://127.0.0.1:1234/Query";
        using (UnityWebRequest www = UnityWebRequest.Post(URL, form))
        {
            //設定資料，並送出
            www.timeout = 5;
            // 若使用 http 協定，需設定 Project Settings/Player/Other/Configuration/Allow downloads over HTTP
            yield return www.SendWebRequest();
            //取得回傳資料
            if (www.isNetworkError || www.isHttpError)
            {
                Debug.LogError($"Error : {www.error}");
            }
            else
            {
                Debug.Log($"Result：{ www.downloadHandler.text}");
            }
        }
    }

}
