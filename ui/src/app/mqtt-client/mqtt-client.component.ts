import { Component, OnDestroy, OnInit, ViewChild } from '@angular/core';
import {animate, state, style, transition, trigger} from '@angular/animations';
import { Subscription } from 'rxjs';
import { IMqttMessage } from "ngx-mqtt";
import { MqttClientService } from '../mqttClient.service';
import { MatTableDataSource } from '@angular/material/table';
import { MatPaginator } from '@angular/material/paginator';
import { MessagePopupComponent } from '../message-popup/message-popup.component';
import { MatDialog } from '@angular/material/dialog';

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

  messages: IMqttMessage[] = [];
  recTS: string[] = [];
  subscription!: Subscription;
  topic: string = '';
  running: boolean = false;
  dataSource!: MatTableDataSource<IMqttMessage>;


  columnsToDisplay = ['Topic', 'Timestamp'];
  columnsToDisplayWithExpand = [...this.columnsToDisplay, 'expand'];
  expandedElement!: IMqttMessage | null;

  @ViewChild(MatPaginator) paginator!: MatPaginator;

  constructor(private mqttClientService: MqttClientService, public dialog: MatDialog,) {}

  ngOnInit(): void {
    this.dataSource = new MatTableDataSource();
  }

  ngOnDestroy(): void {
    
    if (this.subscription) {
      this.subscription.unsubscribe();
    }
    
    this.running = false;
    this.messages = [];
    this.recTS = [];  
}

  subscribeToTopic(topic: string) {
    if (this.topic != '') {
      this.running = true;
      this.subscription = this.mqttClientService.topic(topic).subscribe((data: IMqttMessage) => {
        console.log('Initial time:'+this.getTimestamp());
        data.payload = JSON.parse(data.payload.toString());
        this.messages.push(data);
        this.recTS.push(this.getTimestamp());
        this.dataSource = new MatTableDataSource(this.messages);
        this.dataSource.paginator = this.paginator;
        console.log('Final time:'+this.getTimestamp());
      });
    } else this.dialog.open(MessagePopupComponent, {data: {title: "Error", text: "Topic field cannot be empty!"}});
  }

  unsubscribeTopic() {
    this.subscription.unsubscribe();
    this.running = false; 
  }

  clearData() {
    this.dataSource = new MatTableDataSource();
    this.messages = [];
    this.recTS = [];
  }

  toString(payload: Object): string {
    return JSON.stringify(payload, null, 4);
  }

  getTimestamp(): string {
    var today = new Date();
    var date = today.getFullYear()+'-'+(today.getMonth()+1)+'-'+today.getDate();
    var time = today.getHours() + ":" + today.getMinutes() + ":" + today.getSeconds()+ ":" + today.getMilliseconds();
    var dateTime = date+' '+time;
    return dateTime;
  }
}