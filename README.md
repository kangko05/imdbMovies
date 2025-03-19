### data

This project uses IMDb Non-Commercial Datasets

1. download dataset from [https://dataseets.imdbws.com/](https://datasets.imdbws.com/)
   - name.basics.tsv.gz
   - title.akas.tsv.gz
   - title.basics.tsv.gz
   - title.crew.tsv.gz
   - title.episode.tsv.gz
   - title.principals.tsv.gz
   - title.ratings.tsv.gz
2. store all downloaded '.tsv.gz' in 'data/' folder

Note: Data files are not included in git repository.

### repositry structure

imdbMovies/
├── data/ # IMDb 데이터셋 저장 폴더 (gitignore 처리)
│ ├── .gitkeep # 빈 폴더 유지를 위한 파일
│ └── README.md # 데이터 다운로드 안내
│
├── packages/ # 서비스 폴더
│ ├── movie-service/ # 영화 서비스 (Go/Gin)
│ │ ├── cmd/
│ │ │ └── api/
│ │ │ └── main.go # 서비스 진입점
│ │ ├── internal/
│ │ │ ├── handlers/ # API 엔드포인트 처리
│ │ │ ├── models/ # 데이터 모델
│ │ │ ├── repositories/ # 데이터 액세스
│ │ │ └── services/ # 비즈니스 로직
│ │ ├── pkg/
│ │ │ ├── db/ # 데이터베이스 연결
│ │ │ └── parser/ # TSV 파싱 유틸리티
│ │ ├── go.mod
│ │ └── go.sum
│ │
│ ├── person-service/ # 인물 서비스 (Python/FastAPI)
│ │ ├── app/
│ │ ├── requirements.txt
│ │ └── main.py
│ │
│ ├── rating-service/ # 평점 서비스 (NestJS)
│ │ ├── src/
│ │ └── package.json
│ │
│ ├── user-service/ # 사용자 서비스 (NestJS)
│ │ ├── src/
│ │ └── package.json
│ │
│ └── api-gateway/ # API 게이트웨이 (Express)
│ ├── src/
│ └── package.json
│
├── docker/ # Docker 관련 파일
│ ├── docker-compose.yml
│ └── services/
│ ├── movie-service.dockerfile
│ └── ...
│
├── .gitignore # Git 무시 파일
└── README.md # 프로젝트 설명
