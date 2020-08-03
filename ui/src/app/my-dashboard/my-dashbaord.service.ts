import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { SystemInformation } from './my-dashbaord.model';

@Injectable({
  providedIn: 'root'
})
export class DashboardService {
    constructor(protected httpClient: HttpClient) { }

    getOperatingInformation(): Observable<SystemInformation> {
        return this.httpClient.get<SystemInformation>('/api/v1/os');
    }
}
