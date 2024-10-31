from sklearn.datasets import make_classification
from sklearn.model_selection import train_test_split
from sklearn.ensemble import RandomForestClassifier
import joblib
import os

def train_and_save_model():
    # Create a synthetic dataset
    X, y = make_classification(
        n_samples=1000,
        n_features=4,
        n_informative=2,
        n_redundant=0,
        random_state=42
    )

    # Split the data
    X_train, X_test, y_train, y_test = train_test_split(
        X, y, test_size=0.2, random_state=42
    )

    # Train a simple model
    model = RandomForestClassifier(n_estimators=100, random_state=42)
    model.fit(X_train, y_train)

    # Create models directory if it doesn't exist
    os.makedirs("models", exist_ok=True)

    # Save the model
    joblib.dump(model, "models/model.joblib")
    print("Model trained and saved successfully")

if __name__ == "__main__":
    train_and_save_model()
