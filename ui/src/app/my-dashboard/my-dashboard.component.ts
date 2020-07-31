import { Component, OnInit } from '@angular/core';
import { DashboardService } from './my-dashbaord.service';
import { Observable } from 'rxjs';
import { tap } from 'rxjs/operators';
import { SystemInformation } from './my-dashbaord.model';


@Component({
  selector: 'app-my-dashboard',
  templateUrl: './my-dashboard.component.html',
  styleUrls: ['./my-dashboard.component.scss']
})
export class MyDashboardComponent implements OnInit {

  single = [
    {
      name: 'Free',
      value: 0
    },
    {
      name: 'Assigned',
      value: 0
    }
  ];

  view: any[] = [700, 400];

  gradient = false;
  showLegend = false;
  showLabels = true;
  isDoughnut = false;

  colorScheme = {
    domain: ['#a2f39b', '#ffe485']
  };

  sysInfo: Observable<SystemInformation>;

  constructor(private dashboardService: DashboardService) {}

  ngOnInit() {
    this.sysInfo = this.dashboardService.getOperatingInformation().pipe(
      tap(i => {
        this.single[0].value = i.memory.free
        this.single[1].value = i.memory.total
      })
    );
  }
}
