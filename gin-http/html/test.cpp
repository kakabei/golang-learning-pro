
struct stCodeInfo{

	int ver; 
	int nCode; 
	int createTime; 
	int time; 

	void stCodeInfo()
	{

	}
}; 


std::map<strPhone, CodeInfo> mapCodeInfo; // key : phone, value 验证码信息

int getRandCode()
{
	return 0;
}

std::string FormatPhoneText(stCodeInfo & info)
{

}

int sendToPhone(strPhone, strText)
{
	return 0;
}

// return 1 验证码还在有效期
// return 0 生成并发送成功
// return < 0 失败
int send(const std::string &strPhone, int &nCode)
{
	if (strPhone.empty()) return -1;
	// 1、 检测是否已经存在
	stCodeInfo & oCodeInfo = mapCodeInfo[strPhone];
	time_t nNowTime = Time(); 
	if (oCodeInfo.createTime && oCodeInfo.createTime < nNowTime + 1*60){ // 60 有效
		nCode = oCodeInfo.nCode; 
		return 1; 
	} 

	// 2、生成验证码并保存 

	oCodeInfo.ver = 1;
	oCodeInfo.nCode = getRandCode(); 
	oCodeInfo.createTime = nNowTime; 

	// 3、发送验证码

	std::string strText = FormatPhoneText(oCodeInfo); // 格式化短信内容

	return sendToPhone(strPhone, strText)；  ; 
}

bool check(const & std::string &strPhone ,int nCode)
{
	// 检验验证码
	stCodeInfo &oCodeInfo =  mapCodeInfo[strPhone]; 
	time_t nNowTime = Time(); 
	if (oCodeInfo.createTime && oCodeInfo.createTime > nNowTime + 1*60){
		return -1; // 不存在或过期
	}
	if (oCodeInfo.nCode != nCode){
		return -2 // 验证码错误
	}

	// 删除验证码
	mapCodeInfo.earn(strPhone); 
	return true; 
}