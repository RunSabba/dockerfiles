FROM python:3.12-slim-bullseye

WORKDIR /usr/src/app

COPY requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

COPY flask-app.py .

COPY static ./static

COPY templates ./templates

EXPOSE 5000

CMD ["python", "flask-app.py"]