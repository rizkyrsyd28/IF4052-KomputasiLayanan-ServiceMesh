import os
from fastapi import FastAPI, HTTPException, Request
from fastapi.responses import JSONResponse
from typing import List
import httpx
import logging
from pydantic import BaseModel

app = FastAPI()

class TicketOption(BaseModel):
    id: int
    class_: str
    price: int

    class Config:
        fields = {
            'class_': 'class'
        }

class Response(BaseModel):
    success: bool
    data: List[TicketOption]

PORT = os.getenv("PORT", "4002")
SERVICE_A_HOST = os.getenv("SERVICE_A_HOST", "localhost")
SERVICE_A_PORT = os.getenv("SERVICE_A_PORT", "4000")
API_A_PATH = os.getenv("API_A_PATH", "/api/v1/ticket/JAKARTA")
API_B_PATH = os.getenv("API_B_PATH", "/api/v-1/ticket")
SERVICE_B_HOST = os.getenv("SERVICE_B_HOST", "localhost")
SERVICE_B_PORT = os.getenv("SERVICE_B_PORT", "4001")

@app.get("/api/v1/ticket")
async def get_ticket_options(destination: str, request: Request):
    destination = destination.upper()
    
    async with httpx.AsyncClient() as client:
        try:
            service_a_url = f"http://{SERVICE_A_HOST}:{SERVICE_A_PORT}{request.url.path}?{request.url.query}"
            service_b_url = f"http://{SERVICE_B_HOST}:{SERVICE_B_PORT}{request.url.path}?{request.url.query}"
            
            # Call Service A
            response_a = await client.get(service_a_url)
            response_a.raise_for_status()
            
            # Call Service B
            response_b = await client.get(service_b_url)
            response_b.raise_for_status()
            
            # Combine results
            options_a = response_a.json()["data"]
            options_b = response_b.json()["data"]
            
            # Merge and remove duplicates based on class
            all_options = options_a + options_b
            
            return JSONResponse({
                "success": True,
                "data": all_options
            })
            
        except httpx.RequestError as e:
            raise HTTPException(
                status_code=500,
                detail=f"Error aggregating ticket options"
            )

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=PORT)