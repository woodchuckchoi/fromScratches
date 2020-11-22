# workflow
URI, depth, keyword 갯수를 입력하면 backend는 goroutine 생성
Queue (SQS? RabbitMQ? Redis?)에 작업 queueing
worker(Crawler)는 semaphore를 통해 최대한의 goroutine을 생성하며 데이터 크롤링 후 topic 추출(AWS? 자체 NLP?)
추출한 keyword visualisation

# 프로젝트 목적
go routine을 최대한 많이 사용
go로 service를 만드는 과정을 documentation
monetise할 수 있는 서비스에 대한 brain storming용
