# README
## How to Run:
##### 1. git clone https://jnana_parantapa@bitbucket.org/jnana_parantapa/grab-hack-for-good.git
##### 2. cd grab-hack-for-good
##### 3. Build docker with: `docker build . --tag=nbwc-team-grab-for-good-hackathon`
##### 4. Run docker with: `docker run --publish 9090:9090 nbwc-team-grab-for-good-hackathon`
## Note:
##### 1. For operating the GrabApiOrderDispatcher, you can only use it via back-end, so our back-end will dispatch the order automatically on the cloud.
##### 2. The back-end is already hosted in our cloud, so you can access it via this url: `http://ec2-52-221-244-108.ap-southeast-1.compute.amazonaws.com:9090/api/v1`.
##### 3. We also put API docs (via JSON with Postman format) to make the reading easy.
##### 4. Grab create delivery detail request API on {{url}}/transaction/pay/:transaction-id API on back-end won't work for local build, because it needs AWS credentials to be able queue grab create delivery detail request on AWS Lambda and AWS SQS that need 