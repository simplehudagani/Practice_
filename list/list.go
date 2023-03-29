package list

import (
	"encoding/json"
	"fmt"
	"net/http"
	dbs "sum/connection"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db = dbs.Connection()

type BuzzVyaparProgramBaseline struct {
	Ud                                                             int       `json:"id"`
	ParticipantId                                                  int       `json:"participantId"`
	GfId                                                           int       `json:"gfId"`
	Entry_date                                                     time.Time `json:"entry_date"`
	When_Was_survey_done                                           string    `json:"when_Was_survey_done"`
	Name_of_the_vyapari                                            string    `json:"name_of_the_vyapari"`
	Age                                                            int       `json:"age"`
	Contact_number                                                 int       `json:"contact_number"`
	Village_id                                                     string    `json:"village_id"`
	Location_circle                                                string    `json:"location_circle"`
	Highter_education                                              string    `json:"highter_education"`
	Marital_status                                                 string    `json:"marital_status"`
	Number_of_people_in_the_household                              int       `json:"number_of_people_in_the_household"`
	Do_you_own_a_smart_phone                                       bool      `json:"do_you_own_a_smart_phone"`
	Do_you_have_internet_connection_on_your_smart_phone            bool      `json:"do_you_have_internet_connection_on_your_smart_phone"`
	Are_you_proficient_with_numbers                                bool      `json:"are_you_proficient_with_numbers"`
	Are_you_proficient_with_written_language                       bool      `json:"are_you_proficient_with_written_language"`
	Household_income_monthly                                       int       `json:"household_income_monthly"`
	Over_the_last_month_your_average_income                        int       `json:"over_the_last_month_your_average_income"`
	Your_business_profit_last_month                                int       `json:"your_business_profit_last_month"`
	How_much_monthly_income_would_you_like_to_ideally_earn         int       `json:"how_much_monthly_income_would_you_like_to_ideally_earn"`
	Amount_invested_when_the_business_started                      int       `json:"amount_invested_when_the_business_started"`
	Number_of_years_the_business_has_been_operating                int       `json:"number_of_years_the_business_has_been_operating"`
	Why_do_you_do_business                                         string    `json:"why_do_you_do_business"`
	Tell_us_three_things_about_you_as_an_entrepreneur              string    `json:"tell_us_three_things_about_you_as_an_entrepreneur"`
	Tell_us_three_things_about_your_role_as_a_woman_at_home        string    `json:"tell_us_three_things_about_your_role_as_a_woman_at_home"`
	What_are_your_challenges_in_running_and_growing_your_business  string    `json:"what_are_your_challenges_in_running_and_growing_your_business"`
	What_is_your_plan_to_overcome_these_challenges                 string    `json:"what_is_your_plan_to_overcome_these_challenges"`
	What_are_your_skills                                           string    `json:"what_are_your_skills"`
	What_are_the_resources_available_with_you_for_your_business    string    `json:"what_are_the_resources_available_with_you_for_your_business"`
	Who_is_your_customer_Describe_them_to_us                       string    `json:"who_is_your_customer_Describe_them_to_us"`
	Please_list_down_the_various_components_of_business            string    `json:"please_list_down_the_various_components_of_business"`
	I_know_the_current_state_of_my_business_in_profit_loss_revenue bool      `json:"I_know_the_current_state_of_my_business_in_profit_loss_revenue"`
	What_kind_of_books_of_accounts_do_you_maintain                 string    `json:"what_kind_of_books_of_accounts_do_you_maintain"`
	I_can_generate_ideas_to_solve_my_business_problems             bool      `json:"i_can_generate_ideas_to_solve_my_business_problems"`
	Tell_us_about_one_business_problem                             string    `json:"tell_us_about_one_business_problem"`
	What_is_your_business_goal_Business_impurumenet_madodu         string    `json:"what_is_your_business_goal_Business_impurumenet_madodu"`
	Do_you_have_a_business_plan_to_reach_that_goal                 bool      `json:"do_you_have_a_business_plan_to_reach_that_goal"`
	Can_you_submit_a_business_plan_for_your_goal_to_us_right_now   bool      `json:"can_you_submit_a_business_plan_for_your_goal_to_us_right_now"`
	What_are_the_strenghts_of_your_business                        string    `json:"what_are_the_strenghts_of_your_business"`
	What_are_the_weaknesses_of_your_business                       string    `json:"what_are_the_weaknesses_of_your_business"`
	What_are_the_oppourtunities_for_your_business                  string    `json:"what_are_the_oppourtunities_for_your_business"`
	Are_you_able_to_raise_the_required_finance                     bool      `json:"are_you_able_to_raise_the_required_finance"`
	I_have_taken_a_loan_from                                       string    `json:"i_have_taken_a_loan_from"`
	I_have_trouble_accessing_loan_for_my_business                  bool      `json:"i_have_trouble_accessing_loan_for_my_business"`
	What_are_the_prerequisites_to_access_a_loan                    string    `json:"what_are_the_prerequisites_to_access_a_loan"`
	Taluk_district                                                 string    `json:"taluk_district"`
	Name_of_the_cohort                                             string    `json:"name_of_the_cohort"`
	You_stopped_hold_your_business                                 bool      `json:"you_stopped_hold_your_business"`
	No_hours_engaged_business                                      int       `json:"no_hours_engaged_business"`
	License_for_existing_business                                  bool      `json:"license_for_existing_business"`
	Home_based_work_from_shop                                      string    `json:"home_based_work_from_shop"`
	Loan_currently_availed                                         bool      `json:"loan_currently_availed"`
	Relation_who_borrowed                                          string    `json:"relation_who_borrowed"`
	Loan_total_amount                                              int       `json:"loan_total_amount"`
	Loan_source                                                    string    `json:"loan_source"`
	Loan_repayment_till_date                                       string    `json:"loan_repayment_till_date"`
	Need_additional_skills_business                                bool      `json:"need_additional_skills_business"`
	Skils_what_are_those                                           string    `json:"skils_what_are_those"`
	Module1                                                        bool      `json:"Module1"`
	Module2                                                        bool      `json:"Module2"`
	Module3                                                        bool      `json:"Module3"`
	Module4                                                        bool      `json:"Module4"`
	Module5                                                        bool      `json:"Module5"`
	Module                                                         string    `json:"Module"`
}

type SpoorthiBaselineQuestionnaire struct {
	Id                                                  int    `json:"Id"`
	PartcipantId                                        int    `json:"partcipantId"`
	Email_address                                       string `json:"email_address"`
	GelathiId                                           int    `json:"GelathiId"`
	Entry_date                                          string `json:"entry_date"`
	Spoorthi_Session_Number                             string `json:"Spoorthi_Session_Number"`
	List_down_your_skills                               string `json:"list_down_your_skills"`
	Skills_to_overcome_my_challenges                    string `json:"Skills_to_overcome_my_challenges"`
	Used_skills_resources_combat_challenge              string `json:"used_skills_resources_combat_challenge"`
	Listen_paragraph                                    string `json:"listen_paragraph"`
	Summarize_main_points_paragraph                     string `json:"summarize_main_points_paragraph"`
	Ask_two_questions_help_you_understand               string `json:"ask_two_questions_help_you_understand"`
	Three_infrastructure_of_your_village                string `json:"three_infrastructure_of_your_village"`
	Know_the_need_of_my_community                       string `json:"know_the_need_of_my_community"`
	Together_community_members_community_infrastructure string `json:"together_community_members_community_infrastructure"`
	With_other_community_infrastructure                 string `json:"with_other_community_infrastructure"`
	Bring_someone_together                              string `json:"bring_someone_together"`
	Brought_people_together_incident                    string `json:"brought_people_together_incident"`
	Conflict_with_anyone_ask_position                   string `json:"conflict_with_anyone_ask_position"`
	Conflict_matters_interest_mine                      string `json:"conflict_matters_interest_mine"`
	Module1                                             bool   `json:"module1"`
	Module2                                             bool   `json:"module2"`
	Module3                                             bool   `json:"module3"`
	Module4                                             bool   `json:"module4"`
	Module5                                             bool   `json:"module5"`
	There_puja_at_my_house                              string `json:"there_puja_at_my_house"`
	Module                                              string `json:"Module"`
	FirstName                                           string `json:"firstname"`
}

// const (
// 	DB_HOST     = "buzzwomendatabase-new.cixgcssswxvx.ap-south-1.rds.amazonaws.com"
// 	DB_USER     = "bdms_staff_admin"
// 	DB_PASSWORD = "sfhakjfhyiqundfgs3765827635"
// 	DB_NAME     = "bdms_staff"
// )

//=========================== attendance of spoorthi ======================
func Spoorthiattendace(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,application/json ")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	//func Login(w http.ResponseWriter, r *http.Request) {
	//SetupCORS(&w)
	if r.Method != http.MethodPost {
		w.WriteHeader(405) // Return 405 Method Not Allowed.
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "method not found", "Status Code": "405 "})
		return
	}
	var p SpoorthiBaselineQuestionnaire
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(p.Id)
	if p.Module == "Module1" {
		_, err = db.Exec("update SpoorthiBaselineQuestionnaire set module1=1 where id=?;", p.Id)

		if err != nil {

			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Attendance added successfully for Module1"})
	} else if p.Module == "Module2" {
		_, err = db.Exec("update SpoorthiBaselineQuestionnaire set module2=1 where id=?;", p.Id)

		if err != nil {

			fmt.Println(err)
		}

		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Attendance added successfully for Module2"})
	} else if p.Module == "Module3" {
		_, err = db.Exec("update SpoorthiBaselineQuestionnaire set module3=1 where id=?;", p.Id)

		if err != nil {

			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Attendance added successfully for Module3"})
	} else if p.Module == "Module4" {

		_, err = db.Exec("update SpoorthiBaselineQuestionnaire set module4=1 where id=?;", p.Id)

		if err != nil {

			fmt.Println(err)
		}

		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Attendance added successfully for Module4"})
	} else if p.Module == "Module5" {
		_, err = db.Exec("update SpoorthiBaselineQuestionnaire set module5=1 where id=?;", p.Id)

		if err != nil {

			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Attendance added successfully for Module5"})
	}

}

//============================== green Attendance ============================
func Greenattendance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,application/json ")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	//func Login(w http.ResponseWriter, r *http.Request) {
	//SetupCORS(&w)
	if r.Method != http.MethodPost {
		w.WriteHeader(405) // Return 405 Method Not Allowed.
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "method not found", "Status Code": "405 "})
		return
	}
	var p SpoorthiBaselineQuestionnaire
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(p.Id)
	if p.Module == "Module1" {
		_, err = db.Exec("update GreenBaselineSurvey set module1=1 where id=?;", p.Id)

		if err != nil {

			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Attendance added successfully for Module1"})
	} else if p.Module == "Module2" {

		_, err = db.Exec("update GreenBaselineSurvey set module2=1 where id=?;", p.Id)

		if err != nil {

			fmt.Println(err)
		}

		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Attendance added successfully for Module2"})
	} else if p.Module == "Module3" {
		_, err = db.Exec("update GreenBaselineSurvey set module3=1 where id=?;", p.Id)

		if err != nil {

			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Attendance added successfully for Module3"})
	} else if p.Module == "Module4" {
		_, err = db.Exec("update GreenBaselineSurvey set module4=1 where id=?;", p.Id)

		if err != nil {

			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Attendance added successfully for Module4"})
	} else if p.Module == "Module5" {
		_, err = db.Exec("update GreenBaselineSurvey set module5=1 where id=?;", p.Id)

		if err != nil {

			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Attendance added successfully for Module5"})
	}

}

// ============================= buzz attendance =======================================
func Buzzattendance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,application/json ")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	//func Login(w http.ResponseWriter, r *http.Request) {
	//SetupCORS(&w)
	if r.Method != http.MethodPost {
		w.WriteHeader(405) // Return 405 Method Not Allowed.
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "method not found", "Status Code": "405 "})
		return
	}
	var p SpoorthiBaselineQuestionnaire
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(p.Id)
	if p.Module == "Module1" {
		_, err = db.Exec("update BuzzVyaparProgramBaseline set module1=1 where id=?;", p.Id)

		if err != nil {

			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Attendance added successfully for Module1"})
	} else if p.Module == "Module2" {
		_, err = db.Exec("update BuzzVyaparProgramBaseline set module2=1 where id=?;", p.Id)

		if err != nil {

			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Attendance added successfully for Module2"})
	} else if p.Module == "Module3" {
		_, err = db.Exec("update BuzzVyaparProgramBaseline set module3=1 where id=?;", p.Id)

		if err != nil {

			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Attendance added successfully for Module3"})
	} else if p.Module == "Module4" {
		_, err = db.Exec("update BuzzVyaparProgramBaseline set module4=1 where id=?;", p.Id)

		if err != nil {

			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Attendance added successfully for Module4"})
	} else if p.Module == "Module5" {
		_, err = db.Exec("update BuzzVyaparProgramBaseline set module5=1 where id=?;", p.Id)

		if err != nil {

			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Attendance added successfully for Module5"})
	}

}

//var db *sql.DB

//================================ listing of spoorthi participants based on the Modules1 ===================================

func Spoorthilist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,application/json ")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if r.Method != http.MethodGet {
		w.WriteHeader(405) // Return 405 Method Not Allowed.
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "method not found", "Status Code": "405 "})
		return
	}

	var p SpoorthiBaselineQuestionnaire
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(p.Module1)
	fmt.Println(p.Module)
	if p.Module == "Module1" {

		str := "SELECT firstName FROM training_participants train,SpoorthiBaselineQuestionnaire spoo where (spoo.partcipantId=train.id);"

		rows, err := db.Query(str)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"Message": err, "Status Code": "400 "})
			return
		}

		var FirstName string

		type sporti struct {
			FirstName string
		}

		var result []sporti // creating slice

		for rows.Next() {

			err := rows.Scan(&FirstName)
			if err != nil {
				fmt.Println(err)
			}

			result = append(result, sporti{FirstName: FirstName})

		}
		fmt.Println(result)
		// js, _ := json.Marshal(result)
		// w.Write(js)
		json.NewEncoder(w).Encode(map[string]interface{}{"result": result, "Message": "List all "})

		//result = append(result, SpoorthiBaselineQuestionnaire{Id: Id, partcipantId: partcipantId, email_address: email_address, GelathiId: GelathiId, Spoorthi_Session_Number: Spoorthi_Session_Number, list_down_your_skills: list_down_your_skills, skills_to_overcome_my_challenges: skills_to_overcome_my_challenges, used_skills_resources_combat_challenge: used_skills_resources_combat_challenge, listen_paragraph: listen_paragraph, summarize_main_points_paragraph: summarize_main_points_paragraph, ask_two_questions_help_you_understand: ask_two_questions_help_you_understand, three_infrastructure_of_your_village: three_infrastructure_of_your_village, know_the_need_of_my_community: know_the_need_of_my_community, together_community_members_community_infrastructure: together_community_members_community_infrastructure, with_other_community_infrastructure: with_other_community_infrastructure, bring_someone_together: bring_someone_together, brought_people_together_incident: brought_people_together_incident, conflict_with_anyone_ask_position: conflict_with_anyone_ask_position, conflict_matters_interest_mine: conflict_matters_interest_mine, there_puja_at_my_house: there_puja_at_my_house, module1: module1, module2: module2, module3: module3, module4: module4, module5: module5})
	} else if p.Module == "Module2" {
		str := "SELECT firstName FROM training_participants train,SpoorthiBaselineQuestionnaire spoo where (spoo.partcipantId=train.id) and module1=1;"

		rows, err := db.Query(str)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"Message": err, "Status Code": "400 "})
			return
		}

		var FirstName string

		type sporti struct {
			FirstName string
		}

		var result []sporti // creating slice

		for rows.Next() {

			err := rows.Scan(&FirstName)
			if err != nil {
				fmt.Println(err)
			}

			result = append(result, sporti{FirstName: FirstName})

		}

		fmt.Println(result)
		json.NewEncoder(w).Encode(map[string]interface{}{"result": result, "Message": "List all "})

		//result = append(result, SpoorthiBaselineQuestionnaire{Id: Id, partcipantId: partcipantId, email_address: email_address, GelathiId: GelathiId, Spoorthi_Session_Number: Spoorthi_Session_Number, list_down_your_skills: list_down_your_skills, skills_to_overcome_my_challenges: skills_to_overcome_my_challenges, used_skills_resources_combat_challenge: used_skills_resources_combat_challenge, listen_paragraph: listen_paragraph, summarize_main_points_paragraph: summarize_main_points_paragraph, ask_two_questions_help_you_understand: ask_two_questions_help_you_understand, three_infrastructure_of_your_village: three_infrastructure_of_your_village, know_the_need_of_my_community: know_the_need_of_my_community, together_community_members_community_infrastructure: together_community_members_community_infrastructure, with_other_community_infrastructure: with_other_community_infrastructure, bring_someone_together: bring_someone_together, brought_people_together_incident: brought_people_together_incident, conflict_with_anyone_ask_position: conflict_with_anyone_ask_position, conflict_matters_interest_mine: conflict_matters_interest_mine, there_puja_at_my_house: there_puja_at_my_house, module1: module1, module2: module2, module3: module3, module4: module4, module5: module5})
	} else if p.Module == "Module3" {
		str := "SELECT firstName FROM training_participants train,SpoorthiBaselineQuestionnaire spoo where (spoo.partcipantId=train.id) and module1=1 and module2=1;"

		rows, err := db.Query(str)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"Message": err, "Status Code": "400 "})
			return
		}

		var FirstName string

		type sporti struct {
			FirstName string
		}

		var result []sporti // creating slice

		for rows.Next() {

			err := rows.Scan(&FirstName)
			if err != nil {
				fmt.Println(err)
			}

			result = append(result, sporti{FirstName: FirstName})

		}
		fmt.Println(result)
		json.NewEncoder(w).Encode(map[string]interface{}{"result": result, "Message": "List all "})

		//result = append(result, SpoorthiBaselineQuestionnaire{Id: Id, partcipantId: partcipantId, email_address: email_address, GelathiId: GelathiId, Spoorthi_Session_Number: Spoorthi_Session_Number, list_down_your_skills: list_down_your_skills, skills_to_overcome_my_challenges: skills_to_overcome_my_challenges, used_skills_resources_combat_challenge: used_skills_resources_combat_challenge, listen_paragraph: listen_paragraph, summarize_main_points_paragraph: summarize_main_points_paragraph, ask_two_questions_help_you_understand: ask_two_questions_help_you_understand, three_infrastructure_of_your_village: three_infrastructure_of_your_village, know_the_need_of_my_community: know_the_need_of_my_community, together_community_members_community_infrastructure: together_community_members_community_infrastructure, with_other_community_infrastructure: with_other_community_infrastructure, bring_someone_together: bring_someone_together, brought_people_together_incident: brought_people_together_incident, conflict_with_anyone_ask_position: conflict_with_anyone_ask_position, conflict_matters_interest_mine: conflict_matters_interest_mine, there_puja_at_my_house: there_puja_at_my_house, module1: module1, module2: module2, module3: module3, module4: module4, module5: module5})
	} else if p.Module == "Module4" {
		str := "SELECT firstName FROM training_participants train,SpoorthiBaselineQuestionnaire spoo where (spoo.partcipantId=train.id) and module1=1 and module2=1 and module3=1;"

		rows, err := db.Query(str)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"Message": err, "Status Code": "400 "})
			return
		}

		var FirstName string

		type sporti struct {
			FirstName string
		}

		var result []sporti // creating slice

		for rows.Next() {

			err := rows.Scan(&FirstName)
			if err != nil {
				fmt.Println(err)
			}

			result = append(result, sporti{FirstName: FirstName})

		}
		fmt.Println(result)
		json.NewEncoder(w).Encode(map[string]interface{}{"result": result, "Message": "List all "})

		//result = append(result, SpoorthiBaselineQuestionnaire{Id: Id, partcipantId: partcipantId, email_address: email_address, GelathiId: GelathiId, Spoorthi_Session_Number: Spoorthi_Session_Number, list_down_your_skills: list_down_your_skills, skills_to_overcome_my_challenges: skills_to_overcome_my_challenges, used_skills_resources_combat_challenge: used_skills_resources_combat_challenge, listen_paragraph: listen_paragraph, summarize_main_points_paragraph: summarize_main_points_paragraph, ask_two_questions_help_you_understand: ask_two_questions_help_you_understand, three_infrastructure_of_your_village: three_infrastructure_of_your_village, know_the_need_of_my_community: know_the_need_of_my_community, together_community_members_community_infrastructure: together_community_members_community_infrastructure, with_other_community_infrastructure: with_other_community_infrastructure, bring_someone_together: bring_someone_together, brought_people_together_incident: brought_people_together_incident, conflict_with_anyone_ask_position: conflict_with_anyone_ask_position, conflict_matters_interest_mine: conflict_matters_interest_mine, there_puja_at_my_house: there_puja_at_my_house, module1: module1, module2: module2, module3: module3, module4: module4, module5: module5})
	} else if p.Module == "Module5" {

		str := "SELECT firstName FROM training_participants train,SpoorthiBaselineQuestionnaire spoo where (spoo.partcipantId=train.id) and module1=1 and module2=1 and module3=1 and module4=1;"

		rows, err := db.Query(str)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"Message": err, "Status Code": "400 "})
			return
		}

		var FirstName string

		type sporti struct {
			FirstName string
		}

		var result []sporti // creating slice

		for rows.Next() {

			err := rows.Scan(&FirstName)
			if err != nil {
				fmt.Println(err)
			}

			result = append(result, sporti{FirstName: FirstName})

		}
		fmt.Println(result)
		json.NewEncoder(w).Encode(map[string]interface{}{"result": result, "Message": "List all "})

		//result = append(result, SpoorthiBaselineQuestionnaire{Id: Id, partcipantId: partcipantId, email_address: email_address, GelathiId: GelathiId, Spoorthi_Session_Number: Spoorthi_Session_Number, list_down_your_skills: list_down_your_skills, skills_to_overcome_my_challenges: skills_to_overcome_my_challenges, used_skills_resources_combat_challenge: used_skills_resources_combat_challenge, listen_paragraph: listen_paragraph, summarize_main_points_paragraph: summarize_main_points_paragraph, ask_two_questions_help_you_understand: ask_two_questions_help_you_understand, three_infrastructure_of_your_village: three_infrastructure_of_your_village, know_the_need_of_my_community: know_the_need_of_my_community, together_community_members_community_infrastructure: together_community_members_community_infrastructure, with_other_community_infrastructure: with_other_community_infrastructure, bring_someone_together: bring_someone_together, brought_people_together_incident: brought_people_together_incident, conflict_with_anyone_ask_position: conflict_with_anyone_ask_position, conflict_matters_interest_mine: conflict_matters_interest_mine, there_puja_at_my_house: there_puja_at_my_house, module1: module1, module2: module2, module3: module3, module4: module4, module5: module5})
	} else {
		json.NewEncoder(w).Encode("No modules exits")
		fmt.Println("NoT there")
	}
}

//============================= buzz vyapar ========================
func Buzzlist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,application/json ")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if r.Method != http.MethodPost {
		w.WriteHeader(405) // Return 405 Method Not Allowed.
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "method not found", "Status Code": "405 "})
		return
	}
	//var firstName string
	var p BuzzVyaparProgramBaseline
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		panic(err.Error())
	}

	type sporti struct {
		FirstName string
	}
	//var participantid int
	var firstName string
	if p.Module == "Module1" {
		str := ("select C.firstname from SpoorthiBaselineQuestionnaire A join BuzzVyaparProgramBaseline B on  A.partcipantId = B.partcipantId join training_participants C on C.id=B.partcipantId where A.module5=1;")

		rows, err := db.Query(str)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"Message": err, "Status Code": "400 "})
			return
		}

		var result []sporti // creating slice
		for rows.Next() {
			err1 := rows.Scan(&firstName)
			if err1 != nil {
				fmt.Println(err)
			}

			//var FirstName string

			result = append(result, sporti{FirstName: firstName})

			//fmt.Println(result)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"result": result, "Message": "List all "})

	} else if p.Module == "Module2" {
		str := ("SELECT train.firstName FROM training_participants train,BuzzVyaparProgramBaseline buzz where (buzz.partcipantId=train.id) and module1=1;")
		rows, err := db.Query(str)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"Message": err, "Status Code": "400 "})
			return
		}

		var result []sporti // creating slice
		for rows.Next() {
			err1 := rows.Scan(&firstName)
			if err1 != nil {
				fmt.Println(err)
			}

			//var FirstName string

			result = append(result, sporti{FirstName: firstName})

			//fmt.Println(result)

		}
		json.NewEncoder(w).Encode(map[string]interface{}{"result": result, "Message": "List all "})
	} else if p.Module == "Module3" {
		str := ("SELECT train.firstName FROM training_participants train,BuzzVyaparProgramBaseline buzz where (buzz.partcipantId=train.id) and module1=1 and module2=1;")
		rows, err := db.Query(str)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"Message": err, "Status Code": "400 "})
			return
		}

		var result []sporti // creating slice
		for rows.Next() {
			err1 := rows.Scan(&firstName)
			if err1 != nil {
				fmt.Println(err)
			}

			//var FirstName string

			result = append(result, sporti{FirstName: firstName})

			//fmt.Println(result)

		}
		json.NewEncoder(w).Encode(map[string]interface{}{"result": result, "Message": "List all module3 "})
	} else if p.Module == "Module4" {
		str := ("SELECT train.firstName FROM training_participants train,BuzzVyaparProgramBaseline buzz where (buzz.partcipantId=train.id) and module1=1 and module2=1 and module3=1;")
		rows, err := db.Query(str)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"Message": err, "Status Code": "400 "})
			return
		}

		var result []sporti // creating slice
		for rows.Next() {
			err1 := rows.Scan(&firstName)
			if err1 != nil {
				fmt.Println(err)
			}

			//var FirstName string

			result = append(result, sporti{FirstName: firstName})

			//fmt.Println(result)

		}
		json.NewEncoder(w).Encode(map[string]interface{}{"result": result, "Message": "List all module4 "})
	} else if p.Module == "Module5" {
		str := ("SELECT train.firstName FROM training_participants train,BuzzVyaparProgramBaseline buzz where (buzz.partcipantId=train.id) and module1=1 and module2=1 and module3=1 and module4=1;")
		rows, err := db.Query(str)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"Message": err, "Status Code": "400 "})
			return
		}

		var result []sporti // creating slice
		for rows.Next() {
			err1 := rows.Scan(&firstName)
			if err1 != nil {
				fmt.Println(err)
			}

			//var FirstName string

			result = append(result, sporti{FirstName: firstName})

			//fmt.Println(result)

		}
		json.NewEncoder(w).Encode(map[string]interface{}{"result": result, "Message": "List all module5 "})
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": " No more Modules"})
	}
}

//============================================Green vypar survey ===========================

func Greenlist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,application/json ")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if r.Method != http.MethodPost {
		w.WriteHeader(405) // Return 405 Method Not Allowed.
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": "method not found", "Status Code": "405 "})
		return
	}
	//var firstName string
	var p BuzzVyaparProgramBaseline
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		panic(err.Error())
	}

	type sporti struct {
		FirstName string
	}
	//var participantid int
	var firstName string
	if p.Module == "Module1" {
		str := ("select distinct(C.firstName) from SpoorthiBaselineQuestionnaire A join GreenBaselineSurvey B on  A.partcipantId = B.partcipantId join training_participants C on C.id=B.partcipantId where A.module5=1;")

		rows, err := db.Query(str)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"Message": err, "Status Code": "400 "})
			return
		}

		var result []sporti // creating slice
		for rows.Next() {
			err1 := rows.Scan(&firstName)
			if err1 != nil {
				fmt.Println(err)
			}

			//var FirstName string

			result = append(result, sporti{FirstName: firstName})

			//fmt.Println(result)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"result": result, "Message": "List all "})
	} else if p.Module == "Module2" {
		str := ("SELECT train.firstName FROM training_participants train,GreenBaselineSurvey green where (green.partcipantId=train.id) and module1=1;")

		rows, err := db.Query(str)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"Message": err, "Status Code": "400 "})
			return
		}

		var result []sporti // creating slice
		for rows.Next() {
			err1 := rows.Scan(&firstName)
			if err1 != nil {
				fmt.Println(err)
			}

			//var FirstName string

			result = append(result, sporti{FirstName: firstName})

			//fmt.Println(result)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"result": result, "Message": "List all "})
	} else if p.Module == "Module3" {
		str := ("SELECT train.firstName FROM training_participants train,GreenBaselineSurvey green where (green.partcipantId=train.id) and module1=1 and module2=1;")

		rows, err := db.Query(str)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"Message": err, "Status Code": "400 "})
			return
		}

		var result []sporti // creating slice
		for rows.Next() {
			err1 := rows.Scan(&firstName)
			if err1 != nil {
				fmt.Println(err)
			}

			//var FirstName string

			result = append(result, sporti{FirstName: firstName})

			//fmt.Println(result)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"result": result, "Message": "List all of module3"})
	} else if p.Module == "Module4" {
		str := ("SELECT train.firstName FROM training_participants train,GreenBaselineSurvey green where (green.partcipantId=train.id) and module1=1 and module2=1 and module3=1;")

		rows, err := db.Query(str)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"Message": err, "Status Code": "400 "})
			return
		}

		var result []sporti // creating slice
		for rows.Next() {
			err1 := rows.Scan(&firstName)
			if err1 != nil {
				fmt.Println(err)
			}

			//var FirstName string

			result = append(result, sporti{FirstName: firstName})

			//fmt.Println(result)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"result": result, "Message": "List all of module4"})
	} else if p.Module == "Module5" {
		str := ("SELECT train.firstName FROM training_participants train,GreenBaselineSurvey green where (green.partcipantId=train.id) and module1=1 and module2=1 and module3=1 and module4=1;")

		rows, err := db.Query(str)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"Message": err, "Status Code": "400 "})
			return
		}

		var result []sporti // creating slice
		for rows.Next() {
			err1 := rows.Scan(&firstName)
			if err1 != nil {
				fmt.Println(err)
			}

			//var FirstName string

			result = append(result, sporti{FirstName: firstName})

			//fmt.Println(result)
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"result": result, "Message": "List all of module5"})
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{"Message": " No more Modules"})
	}
}

// 	handler := cors.Default().Handler(mux)
// 	http.ListenAndServe(":5002", handler)
// }
