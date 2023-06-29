# Biometrics Verification API

The Biometrics Verification API is a solution for performing fingerprint matching. It provides a simple and reliable way to verify the identity of individuals by comparing their fingerprint images.

### API Documentation

`POST /fingerprint/`

Paypad:

- `target`: The target fingerprint image file (form-data)
- `sample`: The sample fingerprint image file (form-data)

Response:

The API response will include a JSON object with the following properties:

- `matchScore`: A numerical value representing the similarity score or match confidence.

Example Usage:

```bash
curl -X POST -H "Content-Type: multipart/form-data" -F "sample=@sample_fingerprint.jpg" -F "target=@target_fingerprint.jpg" https://api.example.com/fingerprint/
```

Example Response:

```json
{
  "matchScore": 0.34 // 0.34% match
}
```

Notes:

- The sample and target fingerprint images should be in an appropriate image format (e.g., JPEG) and conform to the API's image requirements.
- The `matchScore` indicates the degree of similarity between the two fingerprints, typically ranging from 0 to 1, where 1 represents a perfect match.

Please refer to the API documentation for further details on authentication, error handling, and any additional functionalities provided by the Biometrics Verification API.
