FROM python:3.9-slim-bullseye AS build

WORKDIR /usr/src/app

COPY requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

COPY . .

#running script to create the model
RUN python train_model.py

FROM python:3.9-slim-bullseye AS app

WORKDIR /app

#our requirements.txt packages reside here
COPY --from=build /usr/local/lib/python3.9/site-packages/ /usr/local/lib/python3.9/site-packages/

COPY --from=build /usr/src/app .

ENV PYTHONBUFFERED=1 \
  PYTHONDONTWRITEBYTECODE=1 \
  PORT=8000

RUN useradd -u 1000 -r -s /bin/false runsabba

RUN chown -R runsabba:runsabba /app

USER runsabba

EXPOSE 8000

CMD ["python", "app.py"]

