package com.kylewbanks.residence.banks.banksresidence;

import android.app.Activity;
import android.os.AsyncTask;
import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.widget.Button;
import android.widget.Toast;
import com.android.volley.Request;
import com.android.volley.RequestQueue;
import com.android.volley.Response;
import com.android.volley.VolleyError;
import com.android.volley.toolbox.StringRequest;
import com.android.volley.toolbox.Volley;

public class BLightActivity extends Activity {

    private static final String TAG = BLightActivity.class.getSimpleName();

    private static final String BASE_URL = "http://192.168.0.200:8080";
    private static final String TOGGLE_ENDPOINT = "/toggle";

    private RequestQueue requestQueue;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_blight);

        requestQueue = Volley.newRequestQueue(this);

        findViewById(R.id.btn_toggle).setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                toggle();
            }
        });
    }

    private void toggle() {
        StringRequest stringRequest = new StringRequest(Request.Method.GET, BASE_URL + TOGGLE_ENDPOINT,
                new Response.Listener<String>() {
                    @Override
                    public void onResponse(String response) {
                        Log.i(TAG, "onResponse: " + response);
                    }
                },
                new Response.ErrorListener() {
                    @Override
                    public void onErrorResponse(VolleyError error) {
                        Toast.makeText(BLightActivity.this, "Something went wrong!", Toast.LENGTH_LONG).show();
                    }
                }
        );

        requestQueue.add(stringRequest);
    }
}
