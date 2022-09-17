import { Component, OnDestroy, OnInit } from '@angular/core';
import {animate, state, style, transition, trigger} from '@angular/animations';
import { Subscription } from 'rxjs';
import { IMqttMessage } from "ngx-mqtt";
import { MqttClientService } from '../mqttClient.service';
import { MatTableDataSource } from '@angular/material/table';

@Component({
  selector: 'app-mqtt-client',
  templateUrl: './mqtt-client.component.html',
  styleUrls: ['./mqtt-client.component.css'],
  animations: [
    trigger('detailExpand', [
      state('collapsed', style({height: '0px', minHeight: '0'})),
      state('expanded', style({height: '*'})),
      transition('expanded <=> collapsed', animate('225ms cubic-bezier(0.4, 0.0, 0.2, 1)')),
    ]),
  ],
})
export class MqttClientComponent implements OnInit, OnDestroy {

  messages: IMqttMessage[] = []
  subscription!: Subscription;
  topic: string = "#"
  dataSource!: MatTableDataSource<IMqttMessage>;

  columnsToDisplay = ['Topic', 'Timestamp'];
  columnsToDisplayWithExpand = [...this.columnsToDisplay, 'expand'];
  expandedElement!: IMqttMessage | null;

  constructor(private mqttClientService: MqttClientService) {}

  ngOnInit(): void {
    this.dataSource = new MatTableDataSource();
    this.subscribeToTopic();
  }

  ngOnDestroy(): void {
    this.subscription.unsubscribe();  
}

  subscribeToTopic() {
    this.subscription = this.mqttClientService.topic(this.topic).subscribe((data: IMqttMessage) => {
      data.payload = JSON.parse(data.payload.toString());
      this.messages.push(data);
      this.dataSource.data = this.messages;
      console.log(this.messages);
      
    });
  }
  toString(payload: Object): string {
    return JSON.stringify(payload, null, 4);
  }

  getTimestamp(): string {
    var today = new Date();
    var date = today.getFullYear()+'-'+(today.getMonth()+1)+'-'+today.getDate();
    var time = today.getHours() + ":" + today.getMinutes() + ":" + today.getSeconds();
    var dateTime = date+' '+time;
    return dateTime;
  }
}